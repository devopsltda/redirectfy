import { CommonModule } from '@angular/common';
import { Component, Input } from '@angular/core';

@Component({
  selector: 'app-button-delete',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './button-delete.component.html',
  styleUrl: './button-delete.component.scss'
})
export class ButtonDeleteComponent {
  @Input() ngClass:string = 'w-9 h-9'
}
