import { CommonModule } from '@angular/common';
import { Component, Input } from '@angular/core';

@Component({
  selector: 'app-icon-box-empty',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './icon-box-empty.component.html',
  styleUrl: './icon-box-empty.component.scss'
})
export class IconBoxEmptyComponent {
  @Input() ngClass:string = "w-64 h-64"
}
