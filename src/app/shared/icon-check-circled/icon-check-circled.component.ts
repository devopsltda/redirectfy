import { CommonModule } from '@angular/common';
import { Component, Input } from '@angular/core';

@Component({
  selector: 'app-icon-check-circled',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './icon-check-circled.component.html',
  styleUrl: './icon-check-circled.component.scss'
})
export class IconCheckCircledComponent {
  @Input() ngClass:string = 'w-24 h-24';
}
