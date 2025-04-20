package mr

import (
	"encoding/json"
	"errors"
	"fmt"
	"hash/fnv"
	"io"
	"io/fs"
	"log"
	"net/rpc"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

type KeyValue struct {
	Key   string
	Value string
}

func ihash(key string) int {
	h := fnv.New32a()
	h.Write([]byte(key))
	return int(h.Sum32() & 0x7fffffff)
}

type ByKey []KeyValue

func (a ByKey) Len() int           { return len(a) }
func (a ByKey) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByKey) Less(i, j int) bool { return a[i].Key < a[j].Key }

func exMTask(filename string, mNumber, nReduce int, mapf func(string, string) []KeyValue) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("cannot open %v", filename)
	}

	content, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("cannot read %v", filename)
	}
	file.Close()
	kva := mapf(filename, string(content))
	mp := map[int]*os.File{}
	for _, kv := range kva {
		rNumber := ihash(kv.Key) % nReduce
		f, ok := mp[rNumber]
		if !ok {
			f, err = os.CreateTemp("", "tmp")
			mp[rNumber] = f
			if !ok {
				if err != nil {
					log.Fatal(err)
				}
			}
		}
		kvj, _ := json.Marshal(kv)
		fmt.Fprintf(f, "%s\n", kvj)
	}

	for rNum, f := range mp {
		os.Rename(f.Name(), fmt.Sprintf("mr-%d-%d", mNumber, rNum))
		f.Close()
	}

}

func WalkDir(root string, rNumber int) ([]string, error) {
	var files []string
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}
		matched, merr := regexp.Match(fmt.Sprintf((`mr-\d-%d`), rNumber), []byte(path))
		if merr != nil {
			return merr
		}

		if matched {
			files = append(files, path)
			return nil
		}
		return nil
	})
	return files, err
}

func exRTask(rNumber int, reducef func(string, []string) string) {
	filenames, _ := WalkDir("./", rNumber)
	data := make([]KeyValue, 0)
	for _, filename := range filenames {
		file, err := os.Open(filename)
		if err != nil {
			log.Fatalf("Cannot open %v. Err: %s", filename, err)
		}
		content, err := io.ReadAll(file)
		if err != nil {
			log.Fatalf("cannot read %v. Err: %s", filename, err)
		}
		file.Close()

		kvstrings := strings.Split(string(content), "\n")
		kv := KeyValue{}

		for _, kvstring := range kvstrings[:len(kvstrings)-1] {
			err := json.Unmarshal([]byte(kvstring), &kv)
			if err != nil {
				log.Fatalf("cannot unmarshal %v. Err: %s", filename, err)
			}
			data = append(data, kv)
		}
		sort.Sort(ByKey(data))

		onname := fmt.Sprintf("mr-out-%d", rNumber)
		ofile, _ := os.Create(onname)

		i := 0
		for i < len(data) {
			j := i + 1
			for j < len(data) && data[j].Key == data[i].Key {
				j++
			}
			values := []string{}
			for k := i; k < j; k++ {
				values = append(values, data[k].Value)
			}
			output := reducef(data[i].Key, values)
			fmt.Fprintf(ofile, "%v %v\n", data[i].Key, output)
			i = j
		}
		ofile.Close()

	}
}

func CallGetTask() (*GetTaskReply, error) {
	args := GetTaskArgs{}
	reply := GetTaskReply{}

	ok := call("Coordinator.GetTask", &args, &reply)
	if ok {
		fmt.Printf("reply.Name '%v', reply.Type '%v'\n ", reply.Name, reply.Type)
		return &reply, nil
	} else {
		return nil, errors.New("call failed")
	}

}

func call(rpcname string, args interface{}, reply interface{}) bool {
	sockname := coordinatorSock()
	c, err := rpc.DialHTTP("unix", sockname)
	if err != nil {
		log.Fatal("dialing:", err)
	}
	defer c.Close()

	err = c.Call(rpcname, args, reply)
	if err == nil {
		return true
	}

	log.Fatalf(err.Error())
	return false
}

func Worker(mapf func(string, string) []KeyValue,
	reducef func(string, []string) string) {
	for {
		reply, err := CallGetTask()
		if err != nil {
			log.Fatal(err)
		}
		if reply.Type == mType {
			exMTask(reply.Name, reply.Number, reply.NReduce, mapf)
			CallUpdateTaskStatus(mType, reply.Name)
		} else {
			exRTask(reply.Number, reducef)
			CallUpdateTaskStatus(rType, reply.Name)
		}
	}
}

func CallUpdateTaskStatus(typ TaskType, name string) error {
	args := MarkTaskAsDoneArgs{
		Name: name,
		Type: typ,
	}
	reply := MarkTaskAsDoneReply{}
	ok := call("Coordinator.UpdateTaskStatus", &args, &reply)
	if ok {
		return nil
	} else {
		return errors.New("call failed")
	}
}
