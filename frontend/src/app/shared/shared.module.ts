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
import { NavbarComponent } from './navbar/navbar.component';
import { IconUserComponent } from './icon-user/icon-user.component';
import { IconSettingsComponent } from './icon-settings/icon-settings.component';
import { IconBoxEmptyComponent } from './icon-box-empty/icon-box-empty.component';
import { ButtonAddComponent } from './button-add/button-add.component';
import { ButtonDeleteComponent } from './button-delete/button-delete.component';
import { ButtonCopyComponent } from './button-copy/button-copy.component';
import { ButtonShareComponent } from './button-share/button-share.component';
import { CardItemLinkComponent } from './card-item-link/card-item-link.component';
import { HeaderComponent } from './header/header.component';
import { ContentWindowComponent } from './content-window/content-window.component';
import { ContentGridComponent } from './content-grid/content-grid.component';
import { ToastModule } from 'primeng/toast';
import { CardSubscriptionPlanComponent } from './card-subscription-plan/card-subscription-plan.component';
import { CardContactMobileComponent } from './card-contact-mobile/card-contact-mobile.component';
import { IconVisaCardComponent } from './icon-visa-card/icon-visa-card.component';
import { RedirectfyPremiumCardComponent } from './redirectfy-premium-card/redirectfy-premium-card.component';
import { RedirectfyPremiumCardDesktopComponent } from './redirectfy-premium-card-desktop/redirectfy-premium-card-desktop.component';
import { IconRightArrowComponent } from './icon-right-arrow/icon-right-arrow.component';


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
    IconVisaCardComponent,
    RedirectfyPremiumCardComponent,
    RedirectfyPremiumCardDesktopComponent,
    IconRightArrowComponent,
    NavbarComponent,
    IconUserComponent,
    IconSettingsComponent,
    IconBoxEmptyComponent,
    ButtonAddComponent,
    ButtonDeleteComponent,
    ButtonCopyComponent,
    ButtonShareComponent,
    CardItemLinkComponent,
    HeaderComponent,
    ContentWindowComponent,
    ContentGridComponent,
    ToastModule,
    CardSubscriptionPlanComponent,
    CardContactMobileComponent

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
    IconMailComponent,
    IconVisaCardComponent,
    RedirectfyPremiumCardComponent,
    RedirectfyPremiumCardDesktopComponent,
    IconRightArrowComponent,
    NavbarComponent,
    IconUserComponent,
    IconSettingsComponent,
    IconBoxEmptyComponent,
    ButtonAddComponent,
    ButtonDeleteComponent,
    ButtonCopyComponent,
    ButtonShareComponent,
    CardItemLinkComponent,
    HeaderComponent,
    ContentWindowComponent,
    ContentGridComponent,
    ToastModule,
    CardSubscriptionPlanComponent,
    CardContactMobileComponent
  ]
})
export class SharedModule { }
