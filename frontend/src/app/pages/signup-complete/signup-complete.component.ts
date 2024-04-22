import { CommonModule } from '@angular/common';
import { Component } from '@angular/core';
import { RouterModule } from '@angular/router';
import { SharedModule } from '../../shared/shared.module';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';

@Component({
  selector: 'app-signup-complete',
  standalone: true,
  imports: [RouterModule,CommonModule,SharedModule,ReactiveFormsModule,FormsModule],
  templateUrl: './signup-complete.component.html',
  styleUrl: './signup-complete.component.scss'
})
export class SignupCompleteComponent {


  onSubmit(){

  }

}
