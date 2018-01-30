import { Component, OnInit, HostBinding } from '@angular/core';
import { OverlayContainer } from '@angular/cdk/overlay';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent implements OnInit {
  title = 'Go Chat';

  @HostBinding('class') componentClass;
  constructor(private overlayContainer: OverlayContainer) {

  }

  ngOnInit() {
    this.overlayContainer.getContainerElement().classList.add('light-theme');
    this.componentClass = 'light-theme';
  }
}
