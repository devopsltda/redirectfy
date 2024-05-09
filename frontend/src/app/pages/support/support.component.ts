import { Component } from '@angular/core';
import { SharedModule } from '../../shared/shared.module';
import { CommonModule } from '@angular/common';
import { RouterModule } from '@angular/router';

@Component({
  selector: 'app-support',
  standalone: true,
  imports: [SharedModule, CommonModule, RouterModule],
  templateUrl: './support.component.html',
  styleUrl: './support.component.scss'
})



export class SupportComponent {
  fileName: string = '';

  onFileSelected(event: any) {
    const file = event.target.files[0];
    if (file) {
      this.fileName = this.truncateFileName(file.name, 25);
    } else {
      this.fileName = '';
    }
  }

  truncateFileName(fileName: string, maxLength: number): string {
    if (fileName.length <= maxLength) {
      return fileName;
    } else {
      const extension = fileName.split('.').pop();
      const truncatedFileName = fileName.substring(0, maxLength - (extension!.length + 5)) + '...' + extension;
      return truncatedFileName;
    }
  }


  
}


