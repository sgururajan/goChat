import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ChatWorkspaceComponent } from './chat-workspace.component';

describe('ChatWorkspaceComponent', () => {
  let component: ChatWorkspaceComponent;
  let fixture: ComponentFixture<ChatWorkspaceComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ChatWorkspaceComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ChatWorkspaceComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
