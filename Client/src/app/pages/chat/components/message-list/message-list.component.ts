import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-message-list',
  templateUrl: './message-list.component.html',
  styleUrls: ['./message-list.component.scss', '../../chat.component.scss']
})
export class MessageListComponent implements OnInit {

  constructor() {
    console.log('MessageListComponent')
  }

  ngOnInit() {
  }

}
