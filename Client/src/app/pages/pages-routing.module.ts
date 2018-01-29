import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { AuthGuard } from '../services/auth.guard';

const routes: Routes = [
  { path: '', redirectTo: 'chat', pathMatch: 'full' },
  {
    path: 'chat', loadChildren: './chat/chat.module#ChatModule',
    canActivate: [AuthGuard], canActivateChild: [AuthGuard], canLoad: [AuthGuard]
  },
  { path: 'auth', loadChildren: './authentication/authentication.module#AuthenticationModule' },
  { path: '**', loadChildren: './page-not-found/page-not-found.module#PageNotFoundModule' }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class PagesRoutingModule { }
