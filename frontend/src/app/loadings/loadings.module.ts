import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { LoadingLoginComponent } from './loading-login/loading-login.component';



@NgModule({
  declarations: [],
  imports: [
    CommonModule,
    LoadingLoginComponent
  ],
  exports:[LoadingLoginComponent]
})
export class LoadingsModule { }
