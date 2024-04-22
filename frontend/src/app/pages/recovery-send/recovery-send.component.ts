import { Component } from '@angular/core';
import { SharedModule } from '../../shared/shared.module';
import { CommonModule } from '@angular/common';
import { RouterModule } from '@angular/router';

@Component({
  selector: 'app-recovery-send',
  standalone: true,
  imports: [SharedModule,CommonModule,RouterModule],
  templateUrl: './recovery-send.component.html',
  styleUrl: './recovery-send.component.scss'
})
export class RecoverySendComponent {

}
