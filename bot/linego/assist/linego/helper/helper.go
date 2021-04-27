package helper

import (
	"../service"
	"../LineThrift"
	"encoding/json"
	"fmt"
	"time"
	"strconv"
)

type mentionMsg struct {
	MENTIONEES []struct {
		S string `json:"S"`
		E string `json:"E"`
		M string `json:"M"`
	} `json:"MENTIONEES"`
}

func MaxRevision (a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func InArray(arr []string, str string) bool {
   for _, a := range arr {
      if a == str {
         return true
      }
   }
   return false
}

func IsAdmin(from string) bool {
	if InArray(service.Creator, from) == true {
		return true
	}
	return false
}

func IsAccess(from string) bool {
	if InArray(service.Creator, from) == true || InArray(service.Squad, from) == true {
		return true
	}
	return false
}

func IsSquad(from string) bool {
	if InArray(service.Squad, from) == true {
		return true
	}
	return false
}

func IsBanned(from string) bool {
	if InArray(service.Banned, from) == true {
		return true
	}
	return false
}

func GetMidFromMentionees(data string) []string{
	var midmen []string
	var midbefore []string
	res := mentionMsg{}
	json.Unmarshal([]byte(data), &res)
	for _, v := range res.MENTIONEES {
		if InArray(midbefore, v.M) == false {
			midbefore = append(midbefore, v.M)
			midmen = append(midmen, v.M)
		} 
	}
	return midmen
}

func Log(optype LineThrift.OpType, logtype string, str string) {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	a:=time.Now().In(loc)
	yyyy := strconv.Itoa(a.Year())
	MM := a.Month().String()
	dd := strconv.Itoa(a.Day())
	hh := a.Hour()
	mm := a.Minute()
	ss := a.Second()
	var hhconv string
	var mmconv string
	var ssconv string
	if hh < 10 {
		hhconv = "0"+strconv.Itoa(hh)
	} else {
		hhconv = strconv.Itoa(hh)
	}
	if mm < 10 {
		mmconv = "0"+strconv.Itoa(mm)
	} else {
		mmconv = strconv.Itoa(mm)
	}
	if ss < 10 {
		ssconv = "0"+strconv.Itoa(ss)
	} else {
		ssconv = strconv.Itoa(ss)
	}
	times := yyyy+"-"+MM+"-"+dd+" "+hhconv+":"+mmconv+":"+ssconv
	fmt.Println("["+times+"]["+optype.String()+"]["+logtype+"]"+str)
}