import { Injectable } from '@angular/core';
import { WebsocketService } from './websocket.service';
import { HttpClient } from '@angular/common/http';
import { Conversation } from '../models/conversationModel';
import { Contact } from '../models/contactsModel';
import { Message } from '../models/messageModel';
import { AuthenticationService } from './authentication.service';
import { Subject } from 'rxjs/Subject';
import 'rxjs/add/operator/map';

@Injectable()
export class ChatService {

  private websocketUrl = '/api/ws';
  private conversationsUrl = '/api/conversations';
  ///api/getmessages/{conversationID}/{skip}/{count}
  private messageListUrl = '/api/getmessages'

  public messages: Subject<Message>;


  constructor(private websocketService: WebsocketService,
    private httpClient: HttpClient,
    private authService: AuthenticationService) {
    this.connectToWebsocket();
  }

  getConversations() {
    this.httpClient.get(this.conversationsUrl).map((res: Conversation[]) => {
      console.log(res);
      return res;
    });
  }

  getMessages(convId: string, currentMsgCount: number, count: number = 10) {
    let url = `${this.messageListUrl}/${convId}/${currentMsgCount}/${count}`;
    this.httpClient.get(url).map((res: Message[]) => {
      console.log(res);
      return res;
    });
  }

  send(conversationId: string, msgBody: string) {
    let authorID = this.authService.getCurrentUserID();
    if (!authorID || authorID.length == 0) return;
    let msg = new Message();
    msg.conversationID = conversationId;
    msg.authorID = authorID;
    msg.body = msgBody;
  }

  private connectToWebsocket() {
    this.messages = <Subject<Message>>this.websocketService.connect(this.websocketUrl).map((res: MessageEvent) => {
      let msg = JSON.parse(res.data) as Message;
      return msg;
    });
  }

}
