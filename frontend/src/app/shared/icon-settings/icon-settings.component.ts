import { CommonModule } from '@angular/common';
import { Component,Input } from '@angular/core';

@Component({
  selector: 'app-icon-settings',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './icon-settings.component.html',
  styleUrl: './icon-settings.component.scss'
})
export class IconSettingsComponent {

  @Input() ngClass!:string
  @Input() width:string = '20'
  @Input() height:string = '20'
  @Input() color:string = '#35B5AE'

}
