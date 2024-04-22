import { CommonModule } from '@angular/common';
import { Component, Input } from '@angular/core';

@Component({
  selector: 'app-icon-redirectify-logo',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './icon-redirectify-logo.component.html',
  styleUrl: './icon-redirectify-logo.component.scss'
})
export class IconRedirectifyLogoComponent {
  @Input() ngClass!:string
  @Input() width:string = '59'
  @Input() height:string = '56'
  @Input() color:string = '#35B5AE'
}
