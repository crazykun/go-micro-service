package registry

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

const ServicePort = ":3000"
const ServiceUrl = "http://localhost" + ServicePort + "/services"

// 服务注册
type registry struct {
	registrations []Registration
	mutex         *sync.Mutex
}

// 添加服务注册
func (r *registry) add(reg Registration) error {
	r.mutex.Lock()
	r.registrations = append(r.registrations, reg)
	r.mutex.Unlock()
	return nil
}

// 取消服务注册
func (r *registry) remove(reg Registration) error {
	for i := range r.registrations {
		if r.registrations[i].ServiceName == reg.ServiceName {
			r.mutex.Lock()
			r.registrations = append(r.registrations[:i], r.registrations[i+1:]...)
			r.mutex.Unlock()
			return nil
		}
	}
	return fmt.Errorf("Service %s not found, Url %s", reg.ServiceName, reg.ServiceUrl)
}

var reg = registry{
	registrations: make([]Registration, 0),
	mutex:         new(sync.Mutex),
}

type RegistryService struct {
}

func (s RegistryService) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	log.Println("Request received")
	switch req.Method {
	case http.MethodPost:
		dec := json.NewDecoder(req.Body)
		var r Registration
		err := dec.Decode(&r)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		log.Printf("Adding service: %v with Url: %s\n", r.ServiceName, r.ServiceUrl)
		err = reg.add(r)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	case http.MethodDelete:
		payload, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		url := string(payload)
		log.Printf("Removing service at URL: %s", url)
		err = reg.remove(Registration{ServiceUrl: url})

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}
