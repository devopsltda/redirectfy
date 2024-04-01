import { Routes } from '@angular/router';
import { LoginComponent } from './pages/login/login.component';
import { SignupComponent } from './pages/signup/signup.component';
import { SignupCompleteComponent } from './pages/signup-complete/signup-complete.component';
import { RecoveryPasswordComponent } from './pages/recovery-password/recovery-password.component';
import { RecoverySendComponent } from './pages/recovery-send/recovery-send.component';
import { ResetPasswordComponent } from './pages/reset-password/reset-password.component';
import { ResetCompleteComponent } from './pages/reset-complete/reset-complete.component';
import { HomeComponent } from './pages/home/home.component';
import { CreateRedirectComponent } from './pages/create-redirect/create-redirect.component';
import path from 'path';
import { Error404Component } from './pages/404/404.component';


export const routes: Routes = [
  {path:'',pathMatch:'full',redirectTo:'login'},
  {path:'login',component:LoginComponent},
  {path:'signup',component:SignupComponent},
  {path:'signup/complete',component:SignupCompleteComponent},
  {path:'recovery',component:RecoveryPasswordComponent},
  {path:'recoverySend',component:RecoverySendComponent},
  {path:'newPassword',component:ResetPasswordComponent},
  {path:'newPasswordComplete',component:ResetCompleteComponent},
  {path:'home',component:HomeComponent},
  {path:'newRedirect',component:CreateRedirectComponent},



  {path:'**',component:Error404Component},
];
