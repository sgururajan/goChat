import { Component, OnInit, ViewChild, HostBinding } from '@angular/core';
import { MatSidenav } from '@angular/material';
import { OverlayContainer } from '@angular/cdk/overlay';

@Component({
  selector: 'app-chat',
  templateUrl: './chat.component.html',
  styleUrls: ['./chat.component.scss']
})
export class ChatComponent implements OnInit {

  @ViewChild('sideNav') sideNav: MatSidenav
  // @HostBinding('class') componentClass;

  selectedList: string = "contacts";

  constructor() { }

  close() {
    this.sideNav.close();
  }

  ngOnInit() {

  }

  onListSelectionChanged(selectedItem: string) {
    console.log('OnListSelectionChanged: ' + selectedItem);
    if (selectedItem && selectedItem.length > 0 && selectedItem != this.selectedList) {
      this.selectedList = selectedItem;
    }
  }

}
