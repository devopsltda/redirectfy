import { Component } from '@angular/core';
import { SharedModule } from 'primeng/api';
import { RouterModule } from '@angular/router';
import { NavbarComponent } from '../../shared/navbar/navbar.component';
import { GridComponent } from '../../shared/grid/grid.component';
import { CommonModule } from '@angular/common';
import { ButtonNotificationComponent } from '../../shared/button-notification/button-notification.component';
import { UserBannerComponent } from '../../shared/user-banner/user-banner.component';

@Component({
  selector: 'app-my-account',
  standalone: true,
  imports: [SharedModule, RouterModule, NavbarComponent, GridComponent, CommonModule, ButtonNotificationComponent, UserBannerComponent],
  templateUrl: './my-account.component.html',
  styleUrl: './my-account.component.scss'
})
export class MyAccountComponent {

}
