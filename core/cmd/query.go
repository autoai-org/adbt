package main

import (
	"github.com/levigross/grequests"
	"github.com/unarxiv/adbt/pkg"
	"log"
)

func clientPost(endpoint string, params map[string]string) *grequests.Response {
	url := "http://127.0.0.1:10591" + endpoint

	resp, err := grequests.Post(url, &grequests.RequestOptions{
		JSON:    params,
		IsAjax:  true,
	})
	if err != nil {
		log.Fatal(err)
	}
	if !resp.Ok {
		log.Fatal("Bad Response from Daemon")
	}
	return resp
}

func clientGet(endpoint string, params map[string]string) *grequests.Response {
	url := "http://127.0.0.1:10591" + endpoint

	resp, err := grequests.Get(url, &grequests.RequestOptions{
		Params:  params,
	})
	if err != nil {
		log.Fatal(err)
	}
	if !resp.Ok {
		log.Fatal("Bad Response from Daemon")
	}
	return resp
}

func clientDelete(endpoint string) *grequests.Response {
	url := "http://127.0.0.1:10591" + endpoint
	resp, err := grequests.Delete(url, &grequests.RequestOptions{})
	if err != nil {
		log.Fatal(err)
	}
	if !resp.Ok {
		log.Fatal("Bad Response from Daemon")
	}
	return resp
}

func createJob(database string, uri string, name string, periods string, time string) bool {
	resp:= clientPost("/jobs",map[string]string{
		"uri": uri,
		"name":   name,
		"database": database,
		"periodicity": periods,
		"time": time,
	})
	return resp.Ok
}

func listJobs() []adbt.Scheduler {
	resp := clientGet("/jobs", map[string]string {

	})
	var respJson []adbt.Scheduler
	resp.JSON(respJson)
	return respJson
}

func deleteJob(identifier string) {

}