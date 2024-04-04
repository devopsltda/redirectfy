import { CommonModule } from '@angular/common';
import { Component, Input } from '@angular/core';

@Component({
  selector: 'app-icon-mail',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './icon-mail.component.html',
  styleUrl: './icon-mail.component.scss'
})
export class IconMailComponent {

  @Input() ngClass:string = 'w-36 h-28'

}
