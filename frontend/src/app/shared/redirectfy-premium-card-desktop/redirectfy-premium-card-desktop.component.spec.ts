import { ComponentFixture, TestBed } from '@angular/core/testing';

import { RedirectfyPremiumCardDesktopComponent } from './redirectfy-premium-card-desktop.component';

describe('RedirectfyPremiumCardDesktopComponent', () => {
  let component: RedirectfyPremiumCardDesktopComponent;
  let fixture: ComponentFixture<RedirectfyPremiumCardDesktopComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [RedirectfyPremiumCardDesktopComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(RedirectfyPremiumCardDesktopComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
