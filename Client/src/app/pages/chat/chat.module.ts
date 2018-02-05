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
import { InitialAvatarDirective } from '../../directives/initial-avatar.directive';
import { DirectivesModule } from '../../directives/directives.module';
import { ContactListComponent } from './components/contact-list/contact-list.component';
import { ConversationComponent } from './components/conversation/conversation.component';

@NgModule({
  imports: [
    CommonModule,
    ChatRoutingModule,
    MaterialModule,
    DirectivesModule
  ],
  declarations: [
    ChatComponent,
    SideNavBarComponent,
    ChatToolbarComponent,
    SideBarComponent,
    MessageListComponent,
    ChatWorkspaceComponent,
    ContactListComponent,
    ConversationComponent
  ]
})
export class ChatModule { }
