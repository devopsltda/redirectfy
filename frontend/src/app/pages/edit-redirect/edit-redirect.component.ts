import { Component, OnInit } from '@angular/core';
import { SharedModule } from '../../shared/shared.module';
import { fadeInOutAnimation } from '../../animations/animations.module';
import { ActivatedRoute } from '@angular/router';
import { FormEditRedirectComponent } from '../../shared/form-edit-redirect/form-edit-redirect.component';
import { RedirectifyApiService } from '../../services/redirectify-api.service';

@Component({
  selector: 'app-edit-redirect',
  standalone: true,
  imports: [SharedModule,FormEditRedirectComponent],
  animations:[fadeInOutAnimation],
  templateUrl: './edit-redirect.component.html',
  styleUrl: './edit-redirect.component.scss'
})
export class EditRedirectComponent implements OnInit {
  hash:string = this.activatedRoute.snapshot.params['hash_redirect']
  redirectData!:any
  constructor(
    private activatedRoute:ActivatedRoute,
    private api:RedirectifyApiService
  ){

  }
async ngOnInit(){
    this.redirectData = await this.api.getRedirect(this.hash)
}

}
