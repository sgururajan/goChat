import { Injectable } from '@angular/core';
import { WebsocketService } from './websocket.service';
import { HttpClient } from '@angular/common/http';
import { Conversation } from '../models/conversationModel';
import { Contact } from '../models/contactsModel';
import { Message } from '../models/messageModel';

@Injectable()
export class ChatService {

  private websocketUrl = '/api/ws';
  private conversationsUrl = '/api/conversations';
  ///api/getmessages/{conversationID}/{skip}/{count}
  private messageListUrl = '/api/getmessages'

  public conversations: Conversation[]
  public contacts: Contact[]
  public messageDict = {};


  constructor(private websocketService: WebsocketService,
    private httpClient: HttpClient) {

  }

  getConversations() {
    this.httpClient.get(this.conversationsUrl).subscribe((res: Conversation[]) => {
      res.forEach(item => {
        let cIndex = this.conversations.findIndex(el => el.conversationID == item.conversationID);
        if (cIndex < 0) {
          this.conversations.push(item);
          if (!this.messageDict.hasOwnProperty(item.conversationID)) {
            this.messageDict[item.conversationID] = new Array<Message>();
          }

          this.getMessages(item.conversationID, 1);
        }
      })
    });
  }

  getMessages(convId: string, count: number = 10) {
    let msgList = this.messageDict[convId] as Message[];
    let currentMsgCount = msgList.length;
    let url = `${this.messageListUrl}/${convId}/${currentMsgCount}/${count}`;
    this.httpClient.get(url).subscribe((res: Message[]) => {
      res.forEach(el => msgList.unshift(el));
    });
  }

  private connectToWebsocket() {
    this.websocketService.connect(this.websocketUrl).subscribe((res: MessageEvent) => {
      let data = res.data;
      let msg = JSON.parse(data) as Message;
      if (msg) {
        let msgList = this.messageDict[msg.conversationID] as Message[];
        //this.messageDict[msg.conversationID].push(msg);
        msgList.unshift(msg);
      }
    });
  }

}
