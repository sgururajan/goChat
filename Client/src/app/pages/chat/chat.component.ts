import { Component, OnInit, ViewChild } from '@angular/core';
import { MatSidenav } from '@angular/material';

@Component({
  selector: 'app-chat',
  templateUrl: './chat.component.html',
  styleUrls: ['./chat.component.scss']
})
export class ChatComponent implements OnInit {

  @ViewChild('sideNav') sideNav: MatSidenav

  constructor() { }

  close() {
    this.sideNav.close();
  }

  ngOnInit() {
  }

}
