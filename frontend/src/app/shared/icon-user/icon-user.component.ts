import { CommonModule } from '@angular/common';
import { Component,Input } from '@angular/core';

@Component({
  selector: 'app-icon-user',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './icon-user.component.html',
  styleUrl: './icon-user.component.scss'
})
export class IconUserComponent {
  @Input() ngClass!:string
  @Input() width:string = '20'
  @Input() height:string = '20'
  @Input() color:string = '#35B5AE'
}
