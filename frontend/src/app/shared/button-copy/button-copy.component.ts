import { CommonModule } from '@angular/common';
import { Component,Input } from '@angular/core';

@Component({
  selector: 'app-button-copy',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './button-copy.component.html',
  styleUrl: './button-copy.component.scss'
})
export class ButtonCopyComponent {
  @Input() ngClass:string = 'w-9 h-9'
}
