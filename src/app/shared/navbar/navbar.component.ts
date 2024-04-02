import { Component } from '@angular/core';
import { SidebarModule } from 'primeng/sidebar';
import { IconRedirectifyLogoComponent } from '../icon-redirectify-logo/icon-redirectify-logo.component';
import { CommonModule } from '@angular/common';
import { IconUserComponent } from '../icon-user/icon-user.component';
import { IconSettingsComponent } from '../icon-settings/icon-settings.component';
import { RedirectifyLogoNameComponent } from '../redirectify-logo-name/redirectify-logo-name.component';
import { Router, RouterLink, RouterModule } from '@angular/router';
import { ButtonNotificationComponent } from '../button-notification/button-notification.component';



@Component({
  selector: 'app-navbar',
  standalone: true,
  imports: [SidebarModule,IconRedirectifyLogoComponent,CommonModule,IconUserComponent,IconSettingsComponent,RedirectifyLogoNameComponent,RouterModule,RouterLink, ButtonNotificationComponent],
  templateUrl: './navbar.component.html',
  styleUrl: './navbar.component.scss'
})
export class NavbarComponent {

  sidebarVisible:boolean = false;
  isHovered:boolean = false;

  toggle(){
    this.sidebarVisible = !this.sidebarVisible
  }
}
