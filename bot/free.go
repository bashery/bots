package main

import (
	"fmt"
	"strings"
	"runtime"
	"time"
	"io/ioutil"
	"os"

	"./linego/LineThrift"
	"./linego/auth"
	"./linego/helper"
	"./linego/service"
	"./linego/talk"
	"./linego/config"
)

var argsRaw = os.Args
var Basename = "//"
var ArgSname = "/"
var name = argsRaw[1]
var AppName = "IOS 8.16.1 iphone7"
//var AppName = argsRaw[2]
var Owner = "u957b70fea5594dc505ced3fd44d23d19"
//var Owner = argsRaw[3]

//type User struct {
//    Group string
//    Link int
//    Invite int
//    Kick int
//}
//var Protection = []User{}
var  ProQR = []string{}
var  ProInvite = []string{}
var  ProKick = []string{}


func checkEqual(list1 []string, list2 []string) bool {
	for _, v := range list1 {
		if helper.InArray(list2, v) {
			return true
		}
	}
	return false
}

func canceling(group string, korban []string) {
					runtime.GOMAXPROCS(10)
					for _, vo := range korban {
						cancel := []string{vo}
						if helper.InArray(service.Banned, vo) {
							go func() {
								talk.CancelGroupInvitation(group, cancel)
							}()
						}
					}
}

func addbl(pelaku string) {
		if !helper.IsBanned(pelaku) && !helper.IsAccess(pelaku) {
				service.Banned = append(service.Banned, pelaku)
		}
}

func addstaff(pelaku string) {
		if !helper.IsCreator(pelaku) {
				service.Creator = append(service.Creator, pelaku)
		}
}
func addstaff2(pelaku string) {
		if !helper.IsStaff(pelaku) {
				service.Admins = append(service.Admins, pelaku)
		}
}
func addsquad(pelaku string) {
		if !helper.IsSquad(pelaku) {
				service.Squad = append(service.Squad, pelaku)
		}
}

func addbots(pelaku string) {
		if !helper.IsBots(pelaku) {
				service.Bots = append(service.Bots, pelaku)
		}
}

func cancelall(group string, korban []string) {
					runtime.GOMAXPROCS(10)
					for _, vo := range korban {
						    cancel := []string{vo}
							go func() {
								talk.CancelGroupInvitation(group, cancel)
							}()
					}
}

func checkurl(group string, pelaku string) {
			runtime.GOMAXPROCS(10)
			go func() {
				res, _ := talk.GetGroup(group)
				cek := res.PreventedJoinByTicket
				if !cek {
					go func() {
						res.PreventedJoinByTicket = true
						talk.UpdateGroup(res)
					}()
					go func() {
						talk.KickoutFromGroup(group, []string{pelaku})
					}()
				}
			}()
}

func restart() {
    procAttr := new(os.ProcAttr)
    procAttr.Files = []*os.File{os.Stdin, os.Stdout, os.Stderr}
    os.StartProcess(os.Args[0], []string{"", "test"}, procAttr)
}
func bot(op *LineThrift.Operation) {
	var Mid string = service.MID

	if op.Type == 26 {
		msg := op.Message
		sender := msg.From_
		var sname = ArgSname
		var rname = Basename
		var txt string
		var pesan = strings.ToLower(msg.Text)

		if strings.HasPrefix(pesan , rname + " ") {
			txt = strings.Replace(pesan, rname + " ", "", 1)
		} else if strings.HasPrefix(pesan , rname) {
			txt = strings.Replace(pesan, rname, "", 1)
		} else if strings.HasPrefix(pesan , sname + " ") {
			txt = strings.Replace(pesan, sname + " ", "", 1)
		} else if strings.HasPrefix(pesan , sname){
			txt = strings.Replace(pesan, sname, "", 1)
		}

		if sender != "" && helper.IsAccess(msg.From_) {
			if strings.Contains(pesan , "kick") {
            	str := fmt.Sprintf("%v",msg.ContentMetadata["MENTION"])
            	taglist := helper.GetMidFromMentionees(str)
            	if taglist != nil {
            		for _,target := range taglist {
            			if !helper.IsBanned(target) {
            				addbl(target)
            			}
            		}
            	}
            }
			if txt == "speed" {
				start := time.Now()
           talk.GetProfile()
				//talk.SendMessage(msg.To, "loading...", 2)
				elapsed := start.Sub(start)
				stringTime := elapsed.String()
				talk.SendMessage(msg.To, stringTime, 2)

			} else if pesan == "rname" {
				//talk.SendMessage(msg.To, rname, 2)
        } else if pesan == "respon" {
				//talk.SendMessage(msg.To, "asyap!!", 2)
        } else if pesan == rname+" help" {
				talk.SendMessage(msg.To, "ᴿᴳᴮᴼᵀˢ ᵛ5\n————————\n\nʜᴇʟᴘ ᴄᴏᴍᴍᴀɴᴅ:\n———————\n1."+rname+" respon\n2."+rname+" ping\n3."+rname+" speed\n4."+rname+" status\n\nᴘʀᴏᴛᴇᴄᴛ:\n———————\n5."+rname+" denyinvite on/off\n6."+rname+" qr on/off\n7."+rname+" clearban\n8."+rname+" clearowner\n9."+rname+" clearsquad\n10."+rname+" clearbots\n11."+rname+" ownerlist\n12."+rname+" adminlist\n13."+rname+" clearchat\n14."+rname+" reboot\n\nOwner Command:\n1."+rname+" upname:\n2."+rname+" sname:\n3."+rname+" rname:", 2)
        } else if pesan == sname+" help" {
				talk.SendMessage(msg.To, "ᴿᴳᴮᴼᵀˢ ᵛ5\n————————\n\nʜᴇʟᴘ ᴄᴏᴍᴍᴀɴᴅ:\n———————\n1."+sname+" respon\n2."+sname+" ping\n3."+sname+" speed\n4."+sname+" status\n\nᴘʀᴏᴛᴇᴄᴛ:\n———————\n5."+sname+" denyinvite on/off\n6."+sname+" qr on/off\n7."+sname+" clearban\n8."+sname+" clearowner\n9."+sname+" clearsquad\n10."+sname+" clearbots\n11."+sname+" ownerlist\n12."+sname+" adminlist\n13."+sname+" clearchat\n14."+sname+" reboot\n\nOwner Command:\n1."+sname+" upname:\n2."+sname+" sname:\n3."+sname+" rname:", 2)
        } else if pesan == "ping" {
				talk.SendMessage(msg.To, "pong", 2)
			} else if pesan == "sname" {
				//talk.SendMessage(msg.To, sname, 2)
			} else if txt == "staff"{
				nm := []string{}
				nmm := []string{}
				for c, a := range service.Creator {
					res,_ := talk.GetContact(a)
					name := res.DisplayName
					c += 0
					name = fmt.Sprintf("%v. %s",c , name)
					nm = append(nm, name)
				}
				stf1 := "Owners:\n"
				str1 := strings.Join(nm, "\n")
				for d, b := range service.Admins {
					res,_ := talk.GetContact(b)
					names := res.DisplayName
					d += 1
					names = fmt.Sprintf("%v. %s",d , names)
					nmm = append(nmm, names)
				}
				stf2 := "\n\nAdmins:\n"
				str2 := strings.Join(nmm, "\n")
				talk.SendMessage(msg.To, stf1+str1+stf2+str2, 2)
        } else if txt == "clearchat" {
                talk.RemoveAllMessages(msg.To)
                talk.SendMessage(msg.To,"Cleared chat message", 2)
        //} else if txt == "reboot" {
                //restart()
                //talk.SendMessage(msg.To,"shutdown..", 2)
            } else if strings.HasPrefix(txt,"sname:") {
            	str := strings.Replace(txt,"sname:","", 1)
            	nm := []string{}
            	nm = append(nm,str)
            	stl := strings.Join(nm,", ")
            	ArgSname = stl
            	talk.SendMessage(msg.To,"Succes update Sname to "+ str,2)
            } else if strings.HasPrefix(txt,"rname:") {
            	str := strings.Replace(txt,"rname:","", 1)
            	nm := []string{}
            	nm = append(nm,str)
            	stl := strings.Join(nm,", ")
            	Basename = stl
            	talk.SendMessage(msg.To,"Succes update Rname to "+ str,2)
            } else if strings.HasPrefix(txt,"upname:") {
            	str := strings.Replace(txt,"upname:","", 1)
            	res,_ := talk.GetProfile()
            	res.DisplayName = str
            	talk.UpdateProfile(res)
            	talk.SendMessage(msg.To,"Succes update Profile Name to "+ str,2)
			} else if txt == "status" {
					anu := talk.InviteIntoGroup(msg.To, service.Creator)
					if anu != nil {
						talk.SendMessage(msg.To, "I'm limit", 2)
					} else {
						talk.SendMessage(msg.To, "I'm ready", 2)
					}
			} else if txt == "banlist"{
				nm := []string{}
				for c, a := range service.Banned {
					res,_ := talk.GetContact(a)
					name := res.DisplayName
					c += 1
					name = fmt.Sprintf("%v. %s",c , name)
					nm = append(nm, name)
				}
				stf := "🔒banlist:\n\n"
				str := strings.Join(nm, "\n")
				talk.SendMessage(msg.To, stf+str, 2)
			} else if txt == "squadlist"{
				nm := []string{}
				for c, a := range service.Squad {
					res,_ := talk.GetContact(a)
					name := res.DisplayName
					c += 1
					name = fmt.Sprintf("%v. %s",c , name)
					nm = append(nm, name)
				}
				stf := "🔒squadlist:\n\n"
				str := strings.Join(nm, "\n")
				talk.SendMessage(msg.To, stf+str, 2)
			} else if txt == "botlist"{
				nm := []string{}
				for c, a := range service.Bots {
					res,_ := talk.GetContact(a)
					name := res.DisplayName
					c += 1
					name = fmt.Sprintf("%v. %s",c , name)
					nm = append(nm, name)
				}
				stf := "Botlist:\n\n"
				str := strings.Join(nm, "\n")
				talk.SendMessage(msg.To, stf+str, 2)
			} else if txt == "bypass" {
				runtime.GOMAXPROCS(100)
				res, _ := talk.GetGroup(msg.To)
				memlist := res.Members
				for _, v := range memlist {
					if !helper.IsAccess(v.Mid) {
						cancel := []string{v.Mid}
						go func() {
							talk.KickoutFromGroup(msg.To, cancel)
						}()
					}
				}
			} else if txt == "cancelall" {
				runtime.GOMAXPROCS(100)
				res, _ := talk.GetGroup(msg.To)
				memlist := res.Invitee
				for _, v := range memlist {
					if !helper.IsAccess(v.Mid) {
						cancel := []string{v.Mid}
						go func() {
							talk.CancelGroupInvitation(msg.To, cancel)
						}()
					}
				}
			} else if txt == "back" {
				talk.LeaveGroup(msg.To)
			} else if txt == "clearban" {
				jum := len(service.Banned)
				service.Banned = []string{}
				str := fmt.Sprintf("Unbaned %v banlist.", jum)
				talk.SendMessage(msg.To, str, 2)
        } else if txt == "clearsquad" {
				jum := len(service.Squad)
				service.Squad = []string{}
				str := fmt.Sprintf("Cleared %v squadlist.", jum)
				talk.SendMessage(msg.To, str, 2)
        } else if txt == "clearbots" {
				jum := len(service.Bots)
				service.Bots = []string{}
				str := fmt.Sprintf("Cleared %v botlist.", jum)
				talk.SendMessage(msg.To, str, 2)
        } else if txt == "clearowner" {
				jum := len(service.Creator)
				service.Creator = []string{}
				str := fmt.Sprintf("Remove %v All Ownerlist.", jum)
				talk.SendMessage(msg.To, str, 2)
        } else if txt == "clearadmin" {
				jum := len(service.Admins)
				service.Admins = []string{}
				str := fmt.Sprintf("Remove %v All Adminlist.", jum)
				talk.SendMessage(msg.To, str, 2)
			} else if txt == "invitesquad" {
				    talk.InviteIntoGroup(msg.To, service.Squad)
			} else if txt == "invitebot" {
				    talk.InviteIntoGroup(msg.To, service.Squad)
			} else if txt == "open" {
				res, _ := talk.GetGroup(msg.To)
				cek := res.PreventedJoinByTicket
				if cek {
					res.PreventedJoinByTicket = false
					talk.UpdateGroup(res)
					fmt.Println("done")
				} else {
					res.PreventedJoinByTicket = true
					talk.UpdateGroup(res)
					fmt.Println("done")
			    }
            } else if strings.HasPrefix(txt , "kick") {
            	str := fmt.Sprintf("%v",msg.ContentMetadata["MENTION"])
            	taglist := helper.GetMidFromMentionees(str)
            	if taglist != nil {
            		for _,target := range taglist {
            			runtime.GOMAXPROCS(10)
            			go func() {
            				talk.KickoutFromGroup(msg.To, []string{target})
            			}()
            		}
            	}
            } else if strings.HasPrefix(txt , "owner:on") {
            	str := fmt.Sprintf("%v",msg.ContentMetadata["MENTION"])
            	taglist := helper.GetMidFromMentionees(str)
            	if taglist != nil {
            		lisa := []string{}
            		for c,target := range taglist {
            			if !helper.IsStaff(target) && !helper.IsCreator(target) {
            				addstaff (target)
            			}
            			res,_ := talk.GetContact(target)
					    name := res.DisplayName
					    c += 1
						name = fmt.Sprintf("%v. %s",c , name)
						lisa = append(lisa, name)
            		}
            		stf := "Owner add\n\n"
					str = strings.Join(lisa, "\n")
					talk.SendMessage(msg.To, stf+str, 2)
            	}

            } else if strings.HasPrefix(txt , "admin:on") {
            	str := fmt.Sprintf("%v",msg.ContentMetadata["MENTION"])
            	taglist := helper.GetMidFromMentionees(str)
            	if taglist != nil {
            		lisa := []string{}
            		for c,target := range taglist {
            			if !helper.IsStaff(target) && !helper.IsCreator(target) {
            				addstaff2 (target)
            			}
            			res,_ := talk.GetContact(target)
					    name := res.DisplayName
					    c += 1
						name = fmt.Sprintf("%v. %s",c , name)
						lisa = append(lisa, name)
            		}
            		stf := "Admins add\n\n"
					str = strings.Join(lisa, "\n")
					talk.SendMessage(msg.To, stf+str, 2)
            	}

            } else if strings.HasPrefix(txt , "squad:on") {
            	str := fmt.Sprintf("%v",msg.ContentMetadata["MENTION"])
            	taglist := helper.GetMidFromMentionees(str)
            	if taglist != nil {
            		lisa := []string{}
            		for c,target := range taglist {
            			if !helper.IsSquad(target) && !helper.IsStaff(target) && !helper.IsCreator(target) {
                       talk.FindAndAddContactsByMid(target)
            				addsquad(target)
            			}
            			res,_ := talk.GetContact(target)
					    name := res.DisplayName
					    c += 1
						name = fmt.Sprintf("%v. %s",c , name)
						lisa = append(lisa, name)
            		}
            		stf := "Added to Squadlist:\n\n"
					str := strings.Join(lisa, "\n")
					talk.SendMessage(msg.To, stf+str, 2)
            	}
            } else if strings.HasPrefix(txt , "bot:on") {
            	str := fmt.Sprintf("%v",msg.ContentMetadata["MENTION"])
            	taglist := helper.GetMidFromMentionees(str)
            	if taglist != nil {
            		lisa := []string{}
            		for c,target := range taglist {
            			if !helper.IsBots(target) && !helper.IsStaff(target) && !helper.IsCreator(target) {
                       talk.FindAndAddContactsByMid(target)
            				addbots(target)
            			}
            			res,_ := talk.GetContact(target)
					    name := res.DisplayName
					    c += 1
						name = fmt.Sprintf("%v. %s",c , name)
						lisa = append(lisa, name)
            		}
            		stf := "Added to Botlist:\n\n"
					str := strings.Join(lisa, "\n")
					talk.SendMessage(msg.To, stf+str, 2)
            	}
            } else if txt == "qr on" {
            	if !helper.InArray(ProQR, msg.To) {
            		ProQR = append(ProQR, msg.To)
            		talk.SendMessage(msg.To, "Link Protection Enabled.", 2)
            	} else {
            		talk.SendMessage(msg.To, "Link Protection Already Enabled.", 2)
            	}
			} else if txt == "denyinvite on" {
            	if !helper.InArray(ProInvite, msg.To) {
            		ProInvite = append(ProInvite, msg.To)
            		talk.SendMessage(msg.To, "Invitation protection enabled.", 2)
            	} else {
            		talk.SendMessage(msg.To, "Invitation protection already Enabled.", 2)
            	}
			} else if txt == "respon" {
            	if !helper.InArray(ProInvite, msg.To) {
            		ProInvite = append(ProInvite, msg.To)
            		talk.SendMessage(msg.To, "Stanby.", 2)
            	} else {
            		talk.SendMessage(msg.To, "Stanby", 2)
            	}
            } else if txt == "protectkick on" {
            	if !helper.InArray(ProKick, msg.To) {
            		ProKick = append(ProKick, msg.To)
            		talk.SendMessage(msg.To, "Kick protection enabled.", 2)
            	} else {
            		talk.SendMessage(msg.To, "Kick protection already Enabled.", 2)
            	}
            } else if txt == "qr off" {
            	if helper.InArray(ProQR, msg.To) {
            		ProQR = helper.Remove(ProQR, msg.To)
            		talk.SendMessage(msg.To, "Link protection disabled.", 2)
            	} else {
            		talk.SendMessage(msg.To, "Link protection already disabled.", 2)
            	}
			} else if txt == "denyinvite off" {
            	if helper.InArray(ProInvite, msg.To) {
            		ProInvite = helper.Remove(ProInvite, msg.To)
            		talk.SendMessage(msg.To, "Invitation protection disabled.", 2)
            	} else {
            		talk.SendMessage(msg.To, "Invitation protection already disabled.", 2)
            	}
			} else if txt == "cekban" {
            	if helper.InArray(ProInvite, msg.To) {
            		ProInvite = helper.Remove(ProInvite, msg.To)
            		talk.SendMessage(msg.To, "Empty", 2)
            	} else {
            		talk.SendMessage(msg.To, "Empty", 2)
            	}
			} else if txt == "prokick off" {
            	if helper.InArray(ProKick, msg.To) {
            		ProQR = helper.Remove(ProKick, msg.To)
            		talk.SendMessage(msg.To, "Kick protection disabled.", 2)
            	} else {
            		talk.SendMessage(msg.To, "Kick protection already disabled.", 2)
            	}
			} else if txt == "set" {
				checking := []string{}
				stf := "Bot Setting:\n\n"
				if helper.InArray(ProKick, msg.To) {
					na := "🔒 Kick Protect  ⇉ On"
					checking = append(checking, na)
				} else {
					na := "🔒 Kick Protect  ⇉ Off"
					checking = append(checking, na)
				}
				if helper.InArray(ProInvite, msg.To) {
					na := "🔒 Deny Invite   ⇉ On"
					checking = append(checking, na)
				} else {
					na := "🔒 Deny Invite   ⇉ Off"
					checking = append(checking, na)
				}
				if helper.InArray(ProQR, msg.To) {
					na := "🔒 Link Protect  ⇉ On"
					checking = append(checking, na)
				} else {
					na := "🔒 Link Protect  ⇉ Off"
					checking = append(checking, na)
				}
				str := strings.Join(checking, "\n")
				talk.SendMessage(msg.To, stf+str, 2)
			}
		}
	} else if op.Type == 19 {
		korban := op.Param3
		kicker := op.Param2
		group := op.Param1

		if korban == Mid {
            go func() {
            	addbl(kicker)
            }()
		} else if helper.InArray(service.Squad, korban) && !helper.IsAccess(kicker) {
			runtime.GOMAXPROCS(10)
			go func() {
				talk.KickoutFromGroup(group, []string{kicker})
			}()
			go func() {
				talk.InviteIntoGroup(group, service.Squad)
			}()
			go func() {
            	addbl(kicker)
            }()
		} else if helper.InArray(service.Bots, korban) && !helper.IsAccess(kicker) {
			runtime.GOMAXPROCS(10)
			go func() {
				talk.KickoutFromGroup(group, []string{kicker})
			}()
			go func() {
				talk.InviteIntoGroup(group, []string{korban})
			}()
			go func() {
            	addbl(kicker)
            }()
		} else if helper.IsAccess(korban) && !helper.IsAccess(kicker) {
			runtime.GOMAXPROCS(10)
			go func() {
				talk.KickoutFromGroup(group, []string{kicker})
			}()
			go func() {
				talk.InviteIntoGroup(group, []string{korban})
			}()
			go func() {
            	addbl(kicker)
            }()
		} else if helper.InArray(ProKick, group) {
			go func(){
				talk.KickoutFromGroup(group, []string{kicker})
			}()
			go func() {
            	addbl(kicker)
            }()
		}

	} else if op.Type == 13 {
		runtime.GOMAXPROCS(30)
	    korban := strings.Split(op.Param3, "\x1e")
		inviter := op.Param2
		group := op.Param1

		if helper.InArray(korban, Mid) {
			if helper.IsAccess(inviter) {
				go func() {
					talk.AcceptGroupInvitation(group)
              //purge(group)
				}()
			}

		} else if helper.InArray(ProInvite, group) {
			if !helper.IsAccess(inviter) {
					go func() {
						cancelall(group, korban)
					}()
					go func() {
						talk.KickoutFromGroup(group, []string{inviter})
					}()
					go func() {
            			addbl(inviter)
           			}()
			}
		} else if checkEqual(korban, service.Banned) {
			if !helper.IsAccess(inviter) {
					go func() {
						canceling(group, korban)
					}()
					go func() {
						talk.KickoutFromGroup(group, []string{inviter})
					}()
					go func() {
            				addbl(inviter)
            		}()
			}
		} else if helper.IsBanned(inviter) && !helper.IsAccess(inviter) {
					go func() {
						cancelall(group, korban)
					}()
					go func() {
						talk.KickoutFromGroup(group, []string{inviter})
					}()
		}

	} else if op.Type == 32 {
		runtime.GOMAXPROCS(10)
		korban := op.Param3
		kicker := op.Param2
		group := op.Param1

		if helper.IsAccess(korban) && !helper.IsAccess(kicker) {
			go func(){
				talk.KickoutFromGroup(group, []string{kicker})
			}()
			go func(){
				talk.InviteIntoGroup(group, service.Squad)
			}()
			go func() {
            	addbl(kicker)
            }()
		} else if helper.InArray(ProKick, group) && !helper.IsAccess(kicker) {
			go func(){
				talk.KickoutFromGroup(group, []string{kicker})
			}()
			go func() {
            	addbl(kicker)
            }()
		}
	} else if op.Type == 17 {
		runtime.GOMAXPROCS(10)
		kicker := op.Param2
		group := op.Param1

		if helper.IsBanned(kicker) && !helper.IsAccess(kicker) {
			go func(){
				talk.KickoutFromGroup(group, []string{kicker})
           
			}()
		}
   } else if op.Type == 16 {
		runtime.GOMAXPROCS(10)
      group := op.Param1
		for _, v := range service.Banned {
			go func(){
           //purge(group)
				talk.KickoutFromGroup(group, []string{v})
			}()
		}
	} else if op.Type == 11 {
		runtime.GOMAXPROCS(10)
		changer := op.Param2
		group := op.Param1

		if helper.InArray(ProQR, group) && !helper.IsAccess(changer) {
			go func(){
				checkurl(group, changer)
			}()
		} else if helper.IsBanned(changer) && !helper.IsAccess(changer) {
			go func(){
				checkurl(group, changer)
			}()
		} 
	}
}


func main() {
    filepath := fmt.Sprintf("/root/golang/token/%s.txt", name)
    b, err := ioutil.ReadFile(filepath)
    if err != nil {
        fmt.Print(err)
    }
    token := string(b)
    config.LINE_APPLICATION = AppName
    service.Squad = append(service.Squad, Owner)
    service.Bots = append(service.Bots, Owner)
	auth.LoginWithAuthToken(token)
	//auth.LoginWithQrCode(true)
	for {
		fetch, _ := talk.FetchOperations(service.Revision, 1)
		if len(fetch) > 0 {
			//if error != nil {
			//	fmt.Println(error)
			//}
			rev := fetch[0].Revision
			service.Revision = helper.MaxRevision(service.Revision, rev)
			bot(fetch[0])
		}
	}
}
