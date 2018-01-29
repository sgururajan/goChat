import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ChatToolbarComponent } from './chat-toolbar.component';

describe('ChatToolbarComponent', () => {
  let component: ChatToolbarComponent;
  let fixture: ComponentFixture<ChatToolbarComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ChatToolbarComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ChatToolbarComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
