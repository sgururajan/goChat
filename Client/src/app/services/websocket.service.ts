import { Injectable } from '@angular/core';
import { Subject } from 'rxjs/Subject';
import { Observable } from 'rxjs/Observable';
import { Observer } from 'rxjs/Observer';

@Injectable()
export class WebsocketService {

  private msgSubject: Subject<MessageEvent>;

  constructor() { }

  public connect(url: string): Subject<MessageEvent> {
    if (!this.msgSubject) {
      this.msgSubject = this.createSocketConnection(url);
      console.log(`Web socket connection created to: ${url}`);
    }

    return this.msgSubject;
  }

  private createSocketConnection(url: string): Subject<MessageEvent> {
    let ws = new WebSocket(url);
    let obs = Observable.create((o: Observer<MessageEvent>) => {
      ws.onmessage = o.next.bind(o);
      ws.onerror = o.error.bind(o);
      ws.onclose = o.complete.bind(o);
      return ws.close.bind(ws)
    });

    let observer = {
      next: (data: object) => {
        if (ws.readyState == WebSocket.OPEN) {
          ws.send(JSON.stringify(data));
        }
      }
    }

    return Subject.create(observer, obs);
  }

}
