# -*- coding: utf-8 -*-
from Bots.line import *
from Bots.akad.ttypes import Message
from Bots.akad.ttypes import ContentType as Type
from Bots.akad.ttypes import ChatRoomAnnouncementContents
from Bots.akad.ttypes import ChatRoomAnnouncement
from Bots.akad.ttypes import Location
from Bots.akad.ttypes import OpType
from anu.thrift.protocol import TCompactProtocol
from anu.thrift.transport import THttpClient
from anu.ttypes import LoginRequest
from datetime import datetime, timedelta
from time import sleep
from bs4 import BeautifulSoup
from humanfriendly import format_timespan, format_size, format_number, format_length
from gtts import gTTS
from threading import Thread, activeCount
from io import StringIO
from multiprocessing import Pool
from googletrans import Translator
from random import choice
from urllib.parse import urlencode
import subprocess as cmd
import time, random, sys, json, base64, subprocess, codecs, threading, LineService, shutil, glob, re, string, os, requests, six, ast, pytz, wikipedia, urllib, urllib.parse, atexit, asyncio, traceback
_session = requests.session()
try:
    import urllib.request as urllib2
except ImportError:
    import urllib2
    
def restart_program():
    python = sys.executable
    os.execl(python, python, * sys.argv)

f = open('token1.txt','r')
token = f.read()

self = LINE("{}".format(str(token)))
f.close()
self.log("Auth Token : " + str(self.authToken))

mid = self.profile.mid
clientMID = self.profile.mid
clientPoll = OEPoll(self)
BotOpen = codecs.open("user.json","r","utf-8")
BotMaker = json.load(BotOpen)

settings = {
  "pict": {}
  }

def headers():
    Headers = {
    'User-Agent': "Line/5.1.2",
    'X-Line-Application': "DESKTOPMAC\t5.1.2\tMAC\t10.9.4-MAVERICKS-x64" ,
    "x-lal": "ja-US_US",
    }
    return Headers
def backupData():
    try:
        backup = BotMaker
        f = codecs.open('user.json','w','utf-8')
        json.dump(backup, f, sort_keys=True, indent=4, ensure_ascii=False)
        return True
    except Exception as error:
        logError(error)
        return False

def logError(text):
    self.log("[ ERROR ] {}".format(str(text)))
    tz = pytz.timezone("Asia/Jakarta")
    timeNow = datetime.now(tz=tz)
    timeHours = datetime.strftime(timeNow,"(%H:%M)")
    day = ["Sunday", "Monday", "Tuesday", "Wednesday", "Thursday","Friday", "Saturday"]
    hari = ["Minggu", "Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu"]
    bulan = ["Januari", "Februari", "Maret", "April", "Mei", "Juni", "Juli", "Agustus", "September", "Oktober", "November", "Desember"]
    inihari = datetime.now(tz=tz)
    hr = inihari.strftime('%A')
    bln = inihari.strftime('%m')
    for i in range(len(day)):
        if hr == day[i]: hasil = hari[i]
    for k in range(0, len(bulan)):
        if bln == str(k): bln = bulan[k-1]
    time = "{}, {} - {} - {} | {}".format(str(hasil), str(inihari.strftime('%d')), str(bln), str(inihari.strftime('%Y')), str(inihari.strftime('%H:%M:%S')))
    with open("logError2.txt","a") as error:
        error.write("\n[ {} ] {}".format(str(time), text))

def sendMentionFooter3(to, text="",ps='', mids=[]):
    arrData = ""
    arr = []
    mention = "@Test bots "
    if mids == []:
        raise Exception("Invalid mids")
    if "@!" in text:
        if text.count("@!") != len(mids):
            raise Exception("Invalid mids")
        texts = text.split("@!")
        textx = ''
        h = ''
        for mid in range(len(mids)):
            h+= str(texts[mid].encode('unicode-escape'))
            textx += str(texts[mid])
            if h != textx:slen = len(textx)+h.count('U0');elen = len(textx)+h.count('U0') + 13
            else:slen = len(textx);elen = len(textx) + 13
            arrData = {'S':str(slen), 'E':str(elen), 'M':mids[mid]}
            arr.append(arrData)
            textx += mention
        textx += str(texts[len(mids)])
    else:
        textx = ''
        slen = len(textx)
        elen = len(textx) + 18
        arrData = {'S':str(slen), 'E':str(elen - 4), 'M':mids[0]}
        arr.append(arrData)
        textx += mention + str(text)
    self.sendMessage(to, textx, {'AGENT_LINK': 'https://line.me/ti/p/~regathb1','AGENT_ICON': "http://dl.profile.line-cdn.net/" + self.getProfile().picturePath,'AGENT_NAME': ps,'MENTION': str('{"MENTIONEES":' + json.dumps(arr) + '}')}, 0)

def bot(op):
    try:
        wizard = BotMaker['wizard']
        admins = BotMaker['admins']
        rname = BotMaker['rname']
        helpms="""
®ᴿᴳ_ᴮᴼᵀˢ ᵛ.3.1🔓

 ⌬ /ping
 ⌬ /abot
 ⌬ /dbot
 ⌬ /login name
 ⌬ /on name
 ⌬ /createbot
 ⌬ /running
 ⌬ /remove
 ⌬ /run for bot
 ⌬ /copy
 ⌬ /kill
 ⌬ /screen
 ⌬ /free
 ⌬ /clear user
 ⌬ /reboot
 ⌬ /leave all
 ⌬ /name
 ⌬ /update pict
 ⌬ /listuser
 ⌬ /clear chace
 ⌬ /loginspam
 ⌬ /loginjs1
 ⌬ /loginjs2
 ⌬ /movefile vps1 to vps2
 ⌬ /zip (nama folder)
 ⌬ /unzip (nama folder)
 ⌬ /booted
 ⌬ /groups
 ⌬ /leave
 ⌬ /remotebypass1
 ⌬ /remotebypass2
 ⌬ /remotekick
 ⌬ /sendspam
 ⌬ /nodeclogin
 ⌬ /rubyrun
 ⌬ /apt install
____________________________
  sᴜᴘᴘᴏʀᴛᴇᴅ ʙʏ: ᴛʜʙ ᴛᴇᴀᴍ
     ᴄᴏᴘʏʀɪɢᴛʜ @2018
   ᴍᴀᴅᴇ ɪɴ ɪɴᴅᴏɴᴇsɪᴀɴ
"""
        if op.type is None:
            pass
        else :
            if op.type == 0:
                pass
            else :
                print("[ {} ] {}".format(str(op.type), OpType._VALUES_TO_NAMES[op.type]))

        if op.type == 26:
            msg = op.message
            msg_id = msg.id
            if msg.contentType == 1:
                if msg._from in wizard:
                    if settings['pict'] == True:
                       xpath = self.downloadObjectMsg(msg.id)
                       self.updateProfilePicture(xpath)
                       self.sendMessage(msg.to, "Picture profile has been update")
                       settings['pict'] = False
          
        if op.type == 13:
            if clientMID in op.param3:
               if op.param2 in wizard or op.param2 in admins:
                self.acceptGroupInvitation(op.param1)
                self.sendMessage(op.param1, "Thanks for invite me")
                
        if op.type in [25,26]:
            try:
                msg = op.message
                pesan = msg.text
                msg_id = msg.id
                txt = msg.text
                receiver = msg.to
                sender = msg._from
                if msg.toType == 0 or msg.toType == 1 or msg.toType == 2:
                    if msg.toType == 0:
                        if sender != self.profile.mid:
                            to = sender
                        else:
                            to = receiver
                    elif msg.toType == 1:
                        to = receiver
                    elif msg.toType == 2:
                        to = receiver
                    if msg.contentType == 0:
                        if pesan is None:
                            return
                        else:    	
                            if sender in wizard:
                                if pesan.lower().startswith("/help"):
                                    self.reply(msg.id,to,helpms)
                                if txt == "/owners":
                                   if wizard == {}:
                                       self.reply(msg.id,to, "The stafflist is empty")
                                   else:
                                       num = 0
                                       mc =""
                                       for mi_d in wizard:
                                           mc += "%i - " % num + self.getContact(mi_d).displayName + "\n"
                                           num = (num+1)
                                       self.reply(msg.id,to, "✡🔯 Test Bot's 🔯✡\n\nOwner List: " + rname +"\n\n👑 Owner 👑\n\n" + mc)
                                       print ("Open Stafflist")
                                if txt.startswith("/delowner"):
                                    if 'MENTION' in msg.contentMetadata.keys() != None:
                                        names = re.findall(r'@(\w+)', msg.text)
                                        mention = ast.literal_eval(msg.contentMetadata['MENTION'])
                                        mentionees = mention['MENTIONEES']
                                        for mention in mentionees:
                                            if mention['M'] in BotMaker['wizard']:
                                                del BotMaker['wizard'][mention['M']]
                                                
                                                self.reply(msg.id,to,self.getContact(mention['M']).displayName +" Success remove in owner from! "+ self.getContact(sender).displayName)
                                            else:
                                                self.reply(msg.id,self.getContact(mention['M']).displayName +" Not in owner")
                                if txt.startswith("/deladmin"):
                                    if 'MENTION' in msg.contentMetadata.keys() != None:
                                        names = re.findall(r'@(\w+)', msg.text)
                                        mention = ast.literal_eval(msg.contentMetadata['MENTION'])
                                        mentionees = mention['MENTIONEES']
                                        for mention in mentionees:
                                            if mention['M'] in BotMaker['admins']:
                                                del BotMaker['admins'][mention['M']]
                                                
                                                self.reply(msg.id,to,self.getContact(mention['M']).displayName +" Success remove in Admin from! "+ self.getContact(sender).displayName)
                                            else:
                                                self.reply(msg.id,self.getContact(mention['M']).displayName +" Not in Admin")

                                if txt.startswith("/owner:on"):
                                    if 'MENTION' in msg.contentMetadata.keys() != None:
                                        names = re.findall(r'@(\w+)', msg.text)
                                        mention = ast.literal_eval(msg.contentMetadata['MENTION'])
                                        mentionees = mention['MENTIONEES']
                                        for mention in mentionees:
                                            if mention['M'] in BotMaker['wizard']:
                                                self.reply(msg.id,to,self.getContact(mention['M']).displayName +" already in owner")
                                            else:
                                                BotMaker['wizard'][mention['M']] = True
                                                self.reply(msg.id,to,self.getContact(mention['M']).displayName +" Success add in owner from! "+ self.getContact(sender).displayName)
                                if txt.startswith("/admin:on"):
                                    if 'MENTION' in msg.contentMetadata.keys() != None:
                                        names = re.findall(r'@(\w+)', msg.text)
                                        mention = ast.literal_eval(msg.contentMetadata['MENTION'])
                                        mentionees = mention['MENTIONEES']
                                        for mention in mentionees:
                                            if mention['M'] in BotMaker['admins']:
                                                self.reply(msg.id,to,self.getContact(mention['M']).displayName +" already in admin")
                                            else:
                                                BotMaker['admins'][mention['M']] = True
                                                self.reply(msg.id,to,self.getContact(mention['M']).displayName +" Success add in admin from! "+ self.getContact(sender).displayName)

                                if pesan.lower().startswith("/back"):
                                    self.leaveGroup(to)
                                if pesan.lower().startswith("./ping"):
                                    self.findAndAddContactsByMid(sender)
                                    self.reply(msg.id,to,"Pong")
                                if pesan.lower().startswith("/login"):
                                   user = str(pesan.split(' ')[1])
                                   a = headers()
                                   a.update({'x-lpqs' : '/api/v4/TalkService.do'})
                                   transport = THttpClient.THttpClient('https://gd2.line.naver.jp/api/v4/TalkService.do')
                                   transport.setCustomHeaders(a)
                                   protocol = TCompactProtocol.TCompactProtocol(transport)
                                   client = LineService.Client(protocol)
                                   qr = client.getAuthQrcode(keepLoggedIn=1, systemName='Test-PC')
                                   link = "line://au/q/" + qr.verifier
                                   if msg.toType == 2:self.reply(msg.id,to, '「 LOGIN Sb 」\nType : GetAuthToken')
                                   else:pass
                                   self.reply(msg.id,to, '「 LOGIN QR 」\nType : GetAuthToken\n> Click link qr for login, only 2 minutes\n{}'.format(link))
                                   a.update({"x-lpqs" : '/api/v4/TalkService.do', 'X-Line-Access': qr.verifier})
                                   json.loads(requests.session().get('https://gd2.line.naver.jp/Q', headers=a).text)
                                   a.update({'x-lpqs' : '/api/v4p/rs'})
                                   transport = THttpClient.THttpClient('https://gd2.line.naver.jp/api/v4p/rs')
                                   transport.setCustomHeaders(a)
                                   protocol = TCompactProtocol.TCompactProtocol(transport)
                                   client = LineService.Client(protocol)
                                   req = LoginRequest()
                                   req.type = 1
                                   req.verifier = qr.verifier
                                   req.e2eeVersion = 1
                                   res = client.loginZ(req)
                                   token = "{}".format(res.authToken)
                                   os.system('cp -r line {}'.format(user))
                                   os.system('cd {} && echo -n "{}" > token.txt'.format(user, token))
                                   os.system("screen -S {} -X quit".format(str(user)))
                                   os.system('screen -dmS {}'.format(user))
                                   os.system('screen -r {} -X stuff "cd {} && python3 login.py \n"'.format(user, user))
                                   self.reply(msg.id,to, 'Done your Sb actived....')
                                if pesan.lower().startswith("/create"):
                                   user = str(pesan.split(' ')[1])
                                   a = headers()
                                   a.update({'x-lpqs' : '/api/v4/TalkService.do'})
                                   transport = THttpClient.THttpClient('https://gd2.line.naver.jp/api/v4/TalkService.do')
                                   transport.setCustomHeaders(a)
                                   protocol = TCompactProtocol.TCompactProtocol(transport)
                                   client = LineService.Client(protocol)
                                   qr = client.getAuthQrcode(keepLoggedIn=1, systemName='Test-PC')
                                   link = "line://au/q/" + qr.verifier
                                   if msg.toType == 2:self.reply(msg.id,to, '「 LOGIN Sb 」\nType : GetAuthToken')
                                   else:pass
                                   self.reply(msg.id,to, '「 LOGIN QR 」\nType : GetAuthToken\n> Click link qr for login, only 2 minutes\n{}'.format(link))
                                   a.update({"x-lpqs" : '/api/v4/TalkService.do', 'X-Line-Access': qr.verifier})
                                   json.loads(requests.session().get('https://gd2.line.naver.jp/Q', headers=a).text)
                                   a.update({'x-lpqs' : '/api/v4p/rs'})
                                   transport = THttpClient.THttpClient('https://gd2.line.naver.jp/api/v4p/rs')
                                   transport.setCustomHeaders(a)
                                   protocol = TCompactProtocol.TCompactProtocol(transport)
                                   client = LineService.Client(protocol)
                                   req = LoginRequest()
                                   req.type = 1
                                   req.verifier = qr.verifier
                                   req.e2eeVersion = 1
                                   res = client.loginZ(req)
                                   token = "{}".format(res.authToken)
                                   os.system('cp -r war {}'.format(user))
                                   os.system('cd {} && echo -n "{}" > token.txt'.format(user, token))
                                   os.system("screen -S {} -X quit".format(str(user)))
                                   os.system('screen -dmS {}'.format(user))
                                   os.system('screen -r {} -X stuff "cd {} && python3 bot.py \n"'.format(user, user))
                                   self.reply(msg.id,to, 'Done your Sb actived....')
                                if pesan.lower().startswith(".abot "):                              	
                                    mid = str(pesan.split(' ')[2])
                                    nama = str(pesan.split(' ')[1])
                                    try:                                    	
                                        BotMaker['myBot'][mid] =  '%s' % nama
                                        self.reply(msg.id,to, "%s has been added to folder" %nama)
                                    except:
                                    	self.reply(msg.id,to,"Eror")

                                if pesan.lower().startswith("/dbot "):
                                   mid = str(pesan.split(' ')[2])
                                   folder = str(pesan.split(' ')[1])                                
                                   try:
                                       del BotMaker['myBot'][mid]
                                       os.system('rm -rf {}'.format(str(folder)))
                                       self.reply(msg.id,to, "folder %s and user has been removed" %folder)
                                   except:
                                       self.reply(msg.id,to, "Not in user folder")
                                       
                                if pesan.lower().startswith("/war "):                                  
                                    user = str(pesan.split(' ')[1])
                                    token = str(pesan.split(' ')[2])
                                    try:
                                        os.system('cp -r war {}'.format(user))
                                        os.system('cd {} && echo -n "{}" > token.txt'.format(user, token))
                                        self.reply(msg.id,to,"Token for %s has been created" %user)
                                    except:
                                        self.reply(msg.id,to,"Error")

                                if pesan.lower().startswith("/on "):
                                    user = pesan.lower().replace('/on ','').split(' ')
                                    for x in user:
                                       try:
                                           os.system("screen -S {} -X quit".format(x))
                                           os.system('screen -dmS {}'.format(x))
                                           os.system('screen -r {} -X stuff "cd {} && python3 login.py \n"'.format(x, x))
                                           time.sleep(3)
                                           self.reply(msg.id,to, "%s has been runned." %x)
                                       except:
                                           self.reply(msg.id,to,"Error")

                                if pesan.lower().startswith("/booted "):
                                    user = pesan.lower().replace('/booted ','').split(' ')
                                    for x in user:
                                       try:
                                           os.system("screen -S {} -X quit".format(x))
                                           os.system('screen -dmS {}'.format(x))
                                           os.system('screen -r {} -X stuff "go run {}.go {}\n"'.format(x, x,x))
                                           #time.sleep(3)
                                           self.reply(msg.id,to, "%s has been runned." %x)
                                       except:
                                           self.reply(msg.id,to,"Error")

                                if pesan.lower().startswith("/cek"):
                                   toke = str(pesan.split(' ')[1])
                                   kk = LINE("{}".format(str(toke)))
                                   G = self.getGroup(msg.to)
                                   ginfo = self.getGroup(msg.to)
                                   G.preventedJoinByTicket = False
                                   self.updateGroup(G)
                                   invsend = 0
                                   Ticket = self.reissueGroupTicket(msg.to)
                                   kk.acceptGroupInvitationByTicket(msg.to,Ticket)
                                   kk.reply(msg.id,to, 'Saya gak Banchat Om...')
                                   time.sleep(1)
                                   try:kk.inviteIntoGroup(to, ["u2237ab0e52df609c9980e97735b9e067"]);has = "Saya jg gak limit Om..."
                                   except:has = "Duh saya limit Om..."
                                   kk.reply(msg.id,to, "{}".format(has)) 


                                if pesan.lower().startswith("/running "):                                  
                                    user = str(pesan.split(' ')[1])
                                    set_token = str(pesan.split(' ')[2])
                                    token = '%s' %str(set_token)
                                    try:
                                        os.system('cp -r free.go {}.go'.format(user))
                                        time.sleep(1)
                                        self.sendMessage(to, "Creating bots....")
                                        os.system("cd token && echo -n '%s' > %s.txt"%(token,user))
                                        os.system("screen -S {} -X quit".format(str(user)))
                                        os.system('screen -dmS {}'.format(user))
                                        os.system('screen -r {} -X stuff "go run {}.go {}\n"'.format(user, user,user,user))
                                        self.reply(msg.id,to, "%s has been runned." %user)
                                    except:
                                        self.reply(msg.id,to,"Error")
                                        
                                if pesan.lower() == "/reboot":
                                   self.reply(msg.id,to, "Rebooting...")
                                   restart_program()
                                   print("REBOOT HELPER")

                                if pesan.lower().startswith("/remove "):
                                   folder = pesan.lower().replace('.remove ','').split(' ')
                                   for x in folder:
                                       os.system('rm -rf {}'.format(x))
                                       time.sleep(2)
                                       self.reply(msg.id,to, "folder %s has been removed" %x)

                                if pesan.lower().startswith("/copy "):
                                   folder = pesan.lower().replace('.copy ','').split(' ')
                                   for x in folder:                                   	
                                      os.system('cp -r bot {}'.format(x))
                                      time.sleep(2)
                                      self.reply(msg.id,to, 'Succes copy folder to name %s'%x)
                
                                if pesan.lower().startswith("/kill "):
                                   screen = pesan.lower().replace('/kill ','').split(' ')
                                   for x in screen:
                                      os.system("screen -S {} -X quit".format(x))
                                      time.sleep(2)
                                      self.reply(msg.id,to, "Screen name %s has been killed" %x)

                                if pesan.lower() == "/listfolder":
                                   if BotMaker["myBot"] == {}:
                                       self.reply(msg.id,to, "List folder empity")
                                   else:                                   
                                       h = [a for a in BotMaker['myBot']]
                                       k = len(h)//100
                                       for aa in range(k+1):     	
                                          msgas = 'List Folder Name:\n'
                                          no=0
                                          for a in h:
                                             no+=1
                                             if BotMaker['myBot'][a] == "":cd = "None."
                                             else:cd = BotMaker['myBot'][a]
                                             if no == len(h):msgas+='\n{}. @!\nFolder  ( {} )'.format(no,cd)
                                             else:msgas+='\n{}. @!\nFolder  ( {} )'.format(no,cd)
                                          msgas += "\n\n®TestBots™"
                                          sendMentionFooter3(to, msgas,' List Folder Name ', h)

                                if pesan.lower() == "/usercontact":
                                   mc = []
                                   for mi_d in BotMaker["myBot"]:
                                     self.reply(msg.id,to, None, contentMetadata={'mid': mi_d}, contentType=13)
                                if pesan.lower() == ".inviteuser":
                                   for i in BotMaker["myBot"]:
                                      self.findAndAddContactsByMid(i)
                                      self.inviteIntoGroup(msg.to, [i])

                                if pesan.startswith("/name "):
                                   xres = pesan.replace("/name ","")
                                   if len(xres) <= 500:
                                      profile = self.getProfile()
                                      profile.displayName = xres
                                      self.updateProfile(profile)
                                      self.reply(msg.id,to, "Name has been update to " + xres)
                            
                                if pesan.lower() == "/update pict":
                                   settings["pict"] = True
                                   self.reply(msg.id,to, "send image")
                            
                                if pesan.lower() == "/leaveall":
                                   gid = self.getGroupIdsJoined()
                                   for i in gid:
                                     self.leaveGroup(i)
                                   else:
                                     self.reply(msg.id,to,"Done")

                                if pesan.lower() == "/free":
                                   process = os.popen('free -m')
                                   a = process.read()
                                   self.reply(msg.id,to, "{}".format(a))
                                   process.close()

                                #if pesan.lower() == "/rebootgo":
                                   #os.system('cd python/golang && go build free.go')
                                   #self.reply(msg.id,to, "Rebooting aplication")

                                if pesan.lower() == "/clear chace":
                                   process = os.popen('sync; echo 3 > /proc/sys/vm/drop_caches')
                                   a = process.read()
                                   self.reply(msg.id,to, "Done")
                                   process.close()

                                if pesan.lower() == "/screen":
                                   proses = os.popen("screen -list")
                                   a = proses.read()
                                   self.reply(msg.id,to, "{}\n®TestBots™".format(str(a)), contentMetadata = {'AGENT_ICON': 'http://dl.profile.line-cdn.net/'+self.getContact(clientMID).pictureStatus, 'AGENT_NAME': '? LIST SCREEN ?', 'AGENT_LINK': 'https://line.me/ti/p/{}'.format(self.getUserTicket().id)})
                                   proses.close()

                                if pesan.lower() == "/clear user":
                                    BotMaker["myBot"] = {}
                                    self.reply(msg.id,to, "All list has been cleared")
                                
                                if "/ti/g/" in pesan.lower():
                                  link_re = re.compile('(?:line\:\/|line\.me\/R)\/ti\/g\/([a-zA-Z0-9_-]+)?')
                                  links = link_re.findall(pesan)
                                  n_links = []
                                  for l in links:
                                    if l not in n_links:
                                      n_links.append(l)
                                  for ticket_id in n_links:
                                     group = self.findGroupByTicket(ticket_id)
                                     self.acceptGroupInvitationByTicket(group.id,ticket_id)

                            if sender in admins:
                                if pesan.lower().startswith("/help"):
                                    self.reply(msg.id,to,helpms)
                                if pesan.lower().startswith("/back"):
                                    self.leaveGroup(to)
                                if pesan.lower().startswith("/cek"):
                                   toke = str(pesan.split(' ')[1])
                                   kk = LINE("{}".format(str(toke)))
                                   G = self.getGroup(msg.to)
                                   ginfo = self.getGroup(msg.to)
                                   G.preventedJoinByTicket = False
                                   self.updateGroup(G)
                                   invsend = 0
                                   Ticket = self.reissueGroupTicket(msg.to)
                                   kk.acceptGroupInvitationByTicket(msg.to,Ticket)
                                   kk.reply(msg.id,to, 'Saya gak Banchat Om...')
                                   time.sleep(1)
                                   try:kk.inviteIntoGroup(to, ["u2237ab0e52df609c9980e97735b9e067"]);has = "Saya jg gak limit Om..."
                                   except:has = "Duh saya limit Om..."
                                   kk.reply(msg.id,to, "{}".format(has))
                                if pesan.lower().startswith("/screen"):
                                    self.reply(msg.id,to,"Percuma gak kan bs cm creator yg bisa")
                                if pesan.lower().startswith("/loginjs1"):
                                    self.reply(msg.id,to,"yeeh ngeyel lu cm creator yg bisa")
                                if pesan.lower().startswith("/loginjs2"):
                                    self.reply(msg.id,to,"dudul lu cm creator yg bisa")
                                if pesan.lower().startswith("/loginspam"):
                                    self.reply(msg.id,to,"ijin dulu ke creator")
                                if pesan.lower().startswith("/sendspam"):
                                    self.reply(msg.id,to,"ijin dulu dudul")
                                if pesan.lower().startswith("/groups"):
                                    self.reply(msg.id,to,"cm creator yg bisa blekok")
                                if pesan.lower().startswith("/running"):
                                    self.reply(msg.id,to,"blekoknyaa dirimuh cm creator yg bisa")                

            except Exception as error:
                logError(error)
                traceback.print_tb(error.__traceback__)



        backupData()
    except Exception as error:
        logError(error)
        traceback.print_tb(error.__traceback__)
                                            
                                        

while True:
    try:
        #delExpire()
        ops = clientPoll.singleTrace(count=50)
        if ops is not None:
            for op in ops:
                bot(op)
                clientPoll.setRevision(op.revision)
    except Exception as error:
        logError(error)
