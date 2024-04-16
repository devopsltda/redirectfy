import { CanActivateFn, Router } from '@angular/router';

import {inject } from '@angular/core';
import { CookieService } from 'ngx-cookie-service';
import { RedirectifyApiService } from '../services/redirectify-api.service';
import { HttpClient, HttpHandler } from '@angular/common/http';


export const authGuard:CanActivateFn =  (route, state) => {

  return inject(CookieService).check('refresh-token')? true : inject(Router).navigate(['/login']);
};
