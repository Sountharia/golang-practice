package main

import "fmt"

func sendMessage(m msg) {
	fmt.Println(m.getMessage());
}

type msg interface{
	getMessage() string
}

type birthdayMsg struct {
	name string
	bday string
}

type sendingReport struct {
	reportName string
	Manager string
	reportType string
}

func (bm birthdayMsg) getMessage() string {
	return fmt.Sprintf("Happy birthday dear %s on %s", bm.name, bm.bday)
}

func (sr sendingReport) getMessage() string {
	return fmt.Sprintf("%s Report to be sent to %s (Manager) in %s format", sr.reportName, sr.Manager, sr.reportType)
}

func test(m msg){
	sendMessage(m)
}

func main(){
	Customer1 := birthdayMsg{
		name: "Sountharia",
		bday: "20th September",
	}
	sendMessage(Customer1)
	test(Customer1)

	sendMessage(sendingReport{
		reportName: "Blind 75",
		Manager: "Sountharia",
		reportType: "PDF",
	})
}
