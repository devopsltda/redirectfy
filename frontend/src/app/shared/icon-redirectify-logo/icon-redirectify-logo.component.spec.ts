import { ComponentFixture, TestBed } from '@angular/core/testing';

import { IconRedirectifyLogoComponent } from './icon-redirectify-logo.component';

describe('IconRedirectifyLogoComponent', () => {
  let component: IconRedirectifyLogoComponent;
  let fixture: ComponentFixture<IconRedirectifyLogoComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [IconRedirectifyLogoComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(IconRedirectifyLogoComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
