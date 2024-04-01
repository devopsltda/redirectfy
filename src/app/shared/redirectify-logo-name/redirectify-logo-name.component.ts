import { Component,Input } from '@angular/core';
import { IconRedirectifyLogoComponent } from '../icon-redirectify-logo/icon-redirectify-logo.component';
import { CommonModule } from '@angular/common';


@Component({
  selector: 'app-redirectify-logo-name',
  standalone: true,
  imports: [IconRedirectifyLogoComponent,CommonModule],
  templateUrl: './redirectify-logo-name.component.html',
  styleUrl: './redirectify-logo-name.component.scss'
})
export class RedirectifyLogoNameComponent {

  @Input() ngClass!:string;

}
