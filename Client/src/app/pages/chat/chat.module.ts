import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { ChatRoutingModule } from './chat-routing.module';
import { ChatComponent } from './chat.component';
import { MaterialModule } from '../../material.module';
import { SideNavBarComponent } from './components/side-nav-bar/side-nav-bar.component';
import { ChatToolbarComponent } from './components/chat-toolbar/chat-toolbar.component';
import { SideBarComponent } from './components/side-bar/side-bar.component';
import { MessageListComponent } from './components/message-list/message-list.component';
import { ChatWorkspaceComponent } from './components/chat-workspace/chat-workspace.component';

@NgModule({
  imports: [
    CommonModule,
    ChatRoutingModule,
    MaterialModule
  ],
  declarations: [ChatComponent, SideNavBarComponent, ChatToolbarComponent, SideBarComponent, MessageListComponent, ChatWorkspaceComponent]
})
export class ChatModule { }
