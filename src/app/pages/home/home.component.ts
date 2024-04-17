
import { Component, OnInit } from '@angular/core';
import { RouterModule } from '@angular/router';
import { SharedModule } from '../../shared/shared.module';
import { fadeInOutAnimation } from '../../animations/animations.module';
import { RedirectifyApiService } from '../../services/redirectify-api.service';
import { PaginatorModule, PaginatorState } from 'primeng/paginator';

interface PageEvent {
  first: number;
  rows: number;
  page: number;
  pageCount: number;
}

@Component({
  selector: 'app-home',
  standalone: true,
  imports: [SharedModule,RouterModule,PaginatorModule],
  animations:[fadeInOutAnimation],
  templateUrl: './home.component.html',
  styleUrl: './home.component.scss'
})
export class HomeComponent implements OnInit {
  first: number  = 0;
  rows: number = 10;

  homeData!:any

  constructor
  (
    private api:RedirectifyApiService,
  )
  {

  }



  onPageChange(event: PaginatorState) {
    this.first = event.first!=undefined?event.first:0
    this.rows = event.rows!=undefined?event.rows:0
}

 async ngOnInit() {
    this.homeData = await this.api.getAllRedirects()
  }

  getDisplayData(){
    const start = this.first;
    const end = start + this.rows;
    return this.homeData.slice(start, end);
  }

}
