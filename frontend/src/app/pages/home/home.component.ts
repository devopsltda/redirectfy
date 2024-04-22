import { Component, OnInit } from '@angular/core';
import { RouterModule } from '@angular/router';
import { SharedModule } from '../../shared/shared.module';
import { fadeInOutAnimation } from '../../animations/animations.module';
import { RedirectifyApiService } from '../../services/redirectify-api.service';
import { PaginatorModule, PaginatorState } from 'primeng/paginator';
import { get } from 'http';

interface PageEvent {
  first: number;
  rows: number;
  page: number;
  pageCount: number;
}

@Component({
  selector: 'app-home',
  standalone: true,
  imports: [SharedModule, RouterModule, PaginatorModule],
  animations: [fadeInOutAnimation],
  templateUrl: './home.component.html',
  styleUrl: './home.component.scss',
})
export class HomeComponent implements OnInit {
  first: number = 0;
  rows: number = 10;

  homeData!: any;

  constructor(private api: RedirectifyApiService) {}
  async ngOnInit() {
    await this.getHomeData()
    console.log(this.homeData);
  }

  onPageChange(event: PaginatorState) {
    this.first = event.first != undefined ? event.first : 0;
    this.rows = event.rows != undefined ? event.rows : 0;
  }

 async cardEvent(event:string,item:any){
    console.log(item)
    if(event == 'delete'){
      try {
        const resApi = await this.api.deleteRedirect(item.codigo_hash)
        this.getHomeData()
      } catch (error) {
        console.log(error)
      }
    }
  }

  async getHomeData(){
    this.homeData = await this.api.getAllRedirects();
  }

  getDisplayData() {
    const start = this.first;
    const end = start + this.rows;
    return this.homeData.slice(start, end);
  }
}
