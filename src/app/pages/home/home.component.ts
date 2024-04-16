
import { Component, OnInit } from '@angular/core';
import { RouterModule } from '@angular/router';
import { SharedModule } from '../../shared/shared.module';
import { fadeInOutAnimation } from '../../animations/animations.module';
import { RedirectifyApiService } from '../../services/redirectify-api.service';



@Component({
  selector: 'app-home',
  standalone: true,
  imports: [SharedModule,RouterModule],
  animations:[fadeInOutAnimation],
  templateUrl: './home.component.html',
  styleUrl: './home.component.scss'
})
export class HomeComponent implements OnInit {

  homeData!:any

  constructor
  (
    private api:RedirectifyApiService,

  )
  {

  }

 async ngOnInit() {
    this.homeData = await this.api.getAllRedirects()
    console.log(this.homeData)
    console.log(await this.api.createRedirect('teste','whatsapp,telegram',[{link:'teste',nome:"teste",plataforma:'teste'}]))
  }

  async getHomeData(){
  }

}
