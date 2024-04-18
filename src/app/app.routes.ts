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
import { Error404Component } from './pages/404/404.component';
import { MyAccountComponent } from './pages/my-account/my-account.component';
import { authGuard } from './guard/auth.guard';
import { BuyNowComponent } from './pages/buy-now/buy-now.component';


export const routes: Routes = [
  {path:'',pathMatch:'full',redirectTo:'login'},
  {path:'login',component:LoginComponent},
  {path:'finishSignup/:hashSignup',component:SignupComponent},
  {path:'signup/complete',component:SignupCompleteComponent},
  {path:'recovery',component:RecoveryPasswordComponent},
  {path:'recoverySend',component:RecoverySendComponent},
  {path:'newPassword',component:ResetPasswordComponent},
  {path:'newPasswordComplete',component:ResetCompleteComponent},
  {path:'home',component:HomeComponent,canActivate:[authGuard]},
  {path:'account',component:MyAccountComponent},
  {path:'newRedirect',component:CreateRedirectComponent},
  {path:'buyNow',component:BuyNowComponent},



  {path:'**',component:Error404Component},
];
