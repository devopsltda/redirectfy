import { CommonModule } from '@angular/common';
import { Component, Input } from '@angular/core';

@Component({
  selector: 'app-icon-telegram',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './icon-telegram.component.html',
  styleUrl: './icon-telegram.component.scss'
})
export class IconTelegramComponent {
  @Input() ngClass:string='w-5 h-5'
}
