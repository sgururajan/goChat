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
  @HostBinding('class') componentClass;

  constructor(private overlayContainer: OverlayContainer) { }

  close() {
    this.sideNav.close();
  }

  ngOnInit() {
    // this.overlayContainer.getContainerElement().classList.add('dark-theme');
    // this.componentClass = 'dark-theme';
  }

}
