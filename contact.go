package main

import "fmt"

type contact struct {
	name              string
	defaulPhoneNumber string
	phoneNumberList   []phoneNumberInfo
}

type phoneNumberInfo struct {
	phoneNumber string
	phoneType   string
}

func (c *contact) call() {
	fmt.Println("Calling " + c.defaulPhoneNumber + " on " + c.name)
}

// golang does not support function overload
func (c *contact) callSpecificPhoneNumber(phone phoneNumberInfo) {
	fmt.Println("Calling " + phone.phoneNumber)
}

func (c *contact) addPhone(phoneType string, phoneNumber string) phoneNumberInfo {
	if c.phoneNumberList == nil {
		c.phoneNumberList = make([]phoneNumberInfo, 0)
	}
	newPhoneNumberInfo := phoneNumberInfo{
		phoneType:   phoneType,
		phoneNumber: phoneNumber,
	}
	c.phoneNumberList = append(c.phoneNumberList, newPhoneNumberInfo)
	return newPhoneNumberInfo
}

func (c *contact) list() {
	fmt.Println("--------------------------")
	fmt.Println("Contact:" + c.name)
	fmt.Printf("Default Phone #: %s\r\n", c.defaulPhoneNumber)
	if len(c.phoneNumberList) <= 0 {
		fmt.Println("No more phone number.")
		return
	}

	for _, v := range c.phoneNumberList {
		fmt.Printf("Phone: %s (%s)\r\n", v.phoneNumber, v.phoneType)
	}
}
