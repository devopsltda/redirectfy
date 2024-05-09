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
import { authGuard, loginGuard } from './guard/auth.guard';
import { BuyNowComponent } from './pages/buy-now/buy-now.component';
import { EditRedirectComponent } from './pages/edit-redirect/edit-redirect.component';
import { SupportComponent } from './pages/support/support.component';
import { RedirecionadorComponent } from './pages/redirecionador/redirecionador.component';
import { ChangePlanComponent } from './pages/change-plan/change-plan.component';


export const routes: Routes = [
  {path:'',pathMatch:'full',redirectTo:'login'},
  {path:'login',component:LoginComponent, canActivate:[loginGuard]},
  {path:'finishSignup/:hashSignup',component:SignupComponent},
  {path:'signup/complete',component:SignupCompleteComponent},
  {path:'recovery',component:RecoveryPasswordComponent},
  {path:'recoverySend',component:RecoverySendComponent},
  {path:'newPassword/:hash',component:ResetPasswordComponent},
  {path:'newPasswordComplete',component:ResetCompleteComponent},
  {path:'account',component:MyAccountComponent,canActivate:[authGuard]},
  {path:'account/changePlan',component:ChangePlanComponent,canActivate:[authGuard]},
  {path:'support',component:SupportComponent,canActivate:[authGuard]},
  {path:'newRedirect',component:CreateRedirectComponent,canActivate:[authGuard]},
  {path:'buyNow',component:BuyNowComponent},
  {path:'home',component:HomeComponent,canActivate:[authGuard]},
  {path:'home/:hash_redirect',component:EditRedirectComponent,canActivate:[authGuard]},
  {path:'to/:hash', component:RedirecionadorComponent},



  {path:'**',component:Error404Component},
];
