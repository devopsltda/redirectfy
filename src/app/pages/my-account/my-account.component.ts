import { Component } from '@angular/core';
import { SharedModule } from 'primeng/api';
import { RouterModule } from '@angular/router';
import { NavbarComponent } from '../../shared/navbar/navbar.component';
import { GridComponent } from '../../shared/grid/grid.component';
import { CommonModule } from '@angular/common';
import { ButtonNotificationComponent } from '../../shared/button-notification/button-notification.component';
import { UserBannerComponent } from '../../shared/user-banner/user-banner.component';
import { RedirectfyPremiumCardComponent } from '../../shared/redirectfy-premium-card/redirectfy-premium-card.component';
import { ButtonPrimaryComponent } from '../../shared/button-primary/button-primary.component';
import { ButtonSecundaryComponent } from '../../shared/button-secundary/button-secundary.component';
import { IconVisaCardComponent } from '../../shared/icon-visa-card/icon-visa-card.component';
import { IconRightArrowComponent } from '../../shared/icon-right-arrow/icon-right-arrow.component';

@Component({
  selector: 'app-my-account',
  standalone: true,
  imports: [SharedModule, RouterModule, NavbarComponent, GridComponent, CommonModule, ButtonNotificationComponent, UserBannerComponent, RedirectfyPremiumCardComponent, ButtonPrimaryComponent, ButtonSecundaryComponent, IconVisaCardComponent, IconRightArrowComponent],
  templateUrl: './my-account.component.html',
  styleUrl: './my-account.component.scss'
})
export class MyAccountComponent {

}
