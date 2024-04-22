import { CommonModule } from '@angular/common';
import { Component,Input } from '@angular/core';

@Component({
  selector: 'app-button-share',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './button-share.component.html',
  styleUrl: './button-share.component.scss'
})
export class ButtonShareComponent {
  @Input() ngClass:string = 'w-9 h-9'
}
