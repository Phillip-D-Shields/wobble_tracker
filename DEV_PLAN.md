# Wobble Tracker - MVP Development Plan

## Milestone 1: Project Foundation & Setup
**Goal:** Get the basic Wails + Svelte + SQLite structure running

### Tasks:
- [ ] Initialize Wails project with Svelte frontend
- [ ] Set up GitHub repository with MIT license
- [ ] Create basic project structure and folders
- [ ] Set up SQLite database connection
- [ ] Create initial database schema migration
- [ ] Add basic README with project description and setup instructions
- [ ] Set up GitHub Projects board with these tasks

**Definition of Done:** App starts up, connects to SQLite, basic "Hello World" interface

---

## Milestone 2: Core Database Schema & Models
**Goal:** Implement all database tables and basic CRUD operations

### Tasks:
- [ ] Create `cats` table schema (id, name, birth_date, diagnosis_date, breed, notes)
- [ ] Create `daily_notes` table schema (id, cat_id, date, note_text, created_at)
- [ ] Create `weekly_entries` table schema (id, cat_id, week_ending, weight, activity_level, meal_quality, med_compliance, playtime_minutes, photo_path)
- [ ] Create `milestones` table schema (id, cat_id, date, milestone_type, description, photo_path)
- [ ] Implement Go structs for each table
- [ ] Create database service layer with basic CRUD methods
- [ ] Write unit tests for database operations

**Definition of Done:** All tables created, basic CRUD operations tested and working

---

## Milestone 3: Cat Profile Management
**Goal:** Users can create, view, edit, and delete cat profiles

### Tasks:
- [ ] Design cat profile creation form UI
- [ ] Implement cat profile creation in backend
- [ ] Create cat profile display/view page
- [ ] Add edit cat profile functionality
- [ ] Create cat selection/switching interface
- [ ] Add basic form validation
- [ ] Handle edge cases (no cats created yet, etc.)

**Definition of Done:** Users can manage multiple cat profiles completely

---

## Milestone 4: Daily Notes Feature
**Goal:** Quick and easy daily note entry system

### Tasks:
- [ ] Design daily notes entry interface
- [ ] Implement daily notes creation endpoint
- [ ] Create notes display/history view
- [ ] Add date picker/calendar navigation
- [ ] Implement edit/delete daily notes
- [ ] Add notes to selected cat context
- [ ] Create "today's notes" quick access

**Definition of Done:** Users can add, view, edit, and delete daily notes for any cat

---

## Milestone 5: Weekly Entries Core
**Goal:** Structured weekly data entry and display

### Tasks:
- [ ] Design weekly entry form UI with all tracking fields
- [ ] Implement weekly entry creation backend
- [ ] Create weekly entry display/view
- [ ] Add week navigation (previous/next week)
- [ ] Implement edit weekly entries
- [ ] Add weekly entry validation
- [ ] Create "current week" auto-detection
- [ ] Handle partial week entries

**Definition of Done:** Users can create and manage comprehensive weekly entries

---

## Milestone 6: Photo Integration
**Goal:** Basic photo upload and storage for weekly entries and milestones

### Tasks:
- [ ] Implement file system photo storage
- [ ] Add photo upload component to weekly entries
- [ ] Create photo display/preview functionality
- [ ] Add photo deletion capability
- [ ] Implement basic photo organization by date
- [ ] Handle photo file size limits and validation
- [ ] Add photo to milestones (basic implementation)

**Definition of Done:** Users can upload, view, and delete photos in weekly entries

---

## Milestone 7: Basic Milestones
**Goal:** Track important events and achievements

### Tasks:
- [ ] Design milestone creation interface
- [ ] Implement milestone types (vet visit, birthday, achievement, other)
- [ ] Create milestone display/timeline view
- [ ] Add milestone editing and deletion
- [ ] Integrate milestones into main dashboard
- [ ] Add milestone photos support
- [ ] Create milestone filtering/searching

**Definition of Done:** Users can track and view important milestones for their cats

---

## Milestone 8: Simple Dashboard & Navigation
**Goal:** Main app interface that ties everything together

### Tasks:
- [ ] Design main dashboard layout
- [ ] Create cat quick-stats display
- [ ] Add recent entries preview
- [ ] Implement app navigation structure
- [ ] Create "today's overview" section
- [ ] Add quick-action buttons (add note, new entry, etc.)
- [ ] Implement responsive design basics
- [ ] Add loading states and error handling

**Definition of Done:** Cohesive app experience with easy navigation between features

---

## Milestone 9: MVP Polish & Testing
**Goal:** Bug-free, user-friendly MVP ready for release

### Tasks:
- [ ] Comprehensive manual testing of all features
- [ ] Fix critical bugs and edge cases
- [ ] Improve error messages and user feedback
- [ ] Add data validation throughout
- [ ] Create app icon and branding
- [ ] Write user documentation/help section
- [ ] Add data export functionality (basic CSV)
- [ ] Performance optimization and cleanup

**Definition of Done:** Stable, polished MVP ready for initial users

---

## Future Considerations (Post-MVP)
- Data visualization and trend charts
- Advanced reporting for vet visits
- Community features and sharing
- Mobile app companion
- Video upload support
- Data backup/restore
- Multi-user support
- Advanced analytics

---

## Recommended GitHub Labels
- `milestone-1`, `milestone-2`, etc.
- `frontend`, `backend`, `database`
- `bug`, `enhancement`, `documentation`
- `good-first-issue`, `help-wanted`
- `priority-high`, `priority-medium`, `priority-low`

## Branch Strategy Suggestion
- `main` - stable releases
- `develop` - integration branch
- Feature branches: `feature/cat-profiles`, `feature/daily-notes`, etc.