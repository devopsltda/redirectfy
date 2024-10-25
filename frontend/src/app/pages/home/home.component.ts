import { Component, OnInit } from '@angular/core';
import { RouterModule } from '@angular/router';
import { SharedModule } from '../../shared/shared.module';
import { fadeInOutAnimation } from '../../animations/animations.module';
import { RedirectifyApiService } from '../../services/redirectify-api.service';
import { PaginatorModule, PaginatorState } from 'primeng/paginator';
import { MessageService } from 'primeng/api';
import { ClipboardService } from 'ngx-clipboard';
import { environment } from '../../../environments/environment';



@Component({
  selector: 'app-home',
  standalone: true,
  imports: [SharedModule, RouterModule, PaginatorModule],
  providers:[MessageService],
  animations: [fadeInOutAnimation],
  templateUrl: './home.component.html',
  styleUrl: './home.component.scss',
})
export class HomeComponent implements OnInit {
  first: number = 0;
  rows: number = 10;

  homeData!: any;

  constructor(private api: RedirectifyApiService,private message:MessageService, private cliboardService:ClipboardService) {}
  async ngOnInit() {
    await this.getHomeData();
  }

  onPageChange(event: PaginatorState) {
    this.first = event.first != undefined ? event.first : 0;
    this.rows = event.rows != undefined ? event.rows : 0;
  }

 async cardEvent(event:string,item:any){
    if(event == 'delete'){
      try {
        const resApi = await this.api.deleteRedirect(item.codigo_hash)
        this.getHomeData()
      } catch (error) {
        this.message.add({summary:"Falha na ação",detail:"Falha ao deletar, ação não concluida",severity:'error'})
      }
    } else if(event == 'copy'){
      this.cliboardService.copyFromContent(`${environment.frontUrl}/to/${item.codigo_hash}`)
      this.message.add({summary:"Link Copiado",detail:"O link de redirecionamento foi copiado com sucesso",severity:'success'})
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
