import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { InitialAvatarDirective } from './initial-avatar.directive';
import { ShowOnMouseOverDirective } from './show-on-mouse-over.directive';

@NgModule({
  imports: [
    CommonModule
  ],
  declarations: [
    InitialAvatarDirective,
    ShowOnMouseOverDirective
  ],
  exports: [
    InitialAvatarDirective,
    ShowOnMouseOverDirective
  ]
})
export class DirectivesModule { }
