import { CanActivateFn, Router } from '@angular/router';

import {inject } from '@angular/core';
import { CookieService } from 'ngx-cookie-service';


export const authGuard:CanActivateFn =  (route, state) => {
  const isAuthenticated =  inject(CookieService).check('access-token') ;
  return isAuthenticated ? true : inject(Router).navigate(['/login']);
};

export const loginGuard:CanActivateFn =  (route, state) => {
  const isAuthenticated =  inject(CookieService).check('access-token') ;
  return isAuthenticated ?  inject(Router).navigate(['/home']) : true;
};
