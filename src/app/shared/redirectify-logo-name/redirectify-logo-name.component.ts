import { Component } from '@angular/core';
import { SharedModule } from '../shared.module';
import { IconRedirectifyLogoComponent } from '../icon-redirectify-logo/icon-redirectify-logo.component';

@Component({
  selector: 'app-redirectify-logo-name',
  standalone: true,
  imports: [IconRedirectifyLogoComponent],
  templateUrl: './redirectify-logo-name.component.html',
  styleUrl: './redirectify-logo-name.component.scss'
})
export class RedirectifyLogoNameComponent {

}
