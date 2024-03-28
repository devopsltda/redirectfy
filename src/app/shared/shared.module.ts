import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { GridComponent } from './grid/grid.component';
import { IconRedirectifyLogoComponent } from './icon-redirectify-logo/icon-redirectify-logo.component';
import { RedirectifyLogoNameComponent } from './redirectify-logo-name/redirectify-logo-name.component';
import { ButtonPrimaryComponent } from './button-primary/button-primary.component';
import { ButtonSecundaryComponent } from './button-secundary/button-secundary.component';
import { InputTextModule } from 'primeng/inputtext';
import { ButtonModule } from 'primeng/button';
import { IconCheckCircledComponent } from './icon-check-circled/icon-check-circled.component';
import { IconMailComponent } from './icon-mail/icon-mail.component';



@NgModule({
  declarations: [],
  imports: [
    GridComponent,
    IconRedirectifyLogoComponent,
    RedirectifyLogoNameComponent,
    ButtonPrimaryComponent,
    ButtonSecundaryComponent,
    InputTextModule,
    ButtonModule,
    IconCheckCircledComponent,
    IconMailComponent,
  ],
  exports:[
    GridComponent,
    IconRedirectifyLogoComponent,
    RedirectifyLogoNameComponent,
    ButtonPrimaryComponent,
    ButtonSecundaryComponent,
    InputTextModule,
    ButtonModule,
    IconCheckCircledComponent,
    IconMailComponent
  ]
})
export class SharedModule { }
