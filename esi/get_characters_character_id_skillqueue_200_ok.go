/* 
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * OpenAPI spec version: 0.3.1.dev2
 * 
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package esi

import (
	"time"
)

// 200 ok object
type GetCharactersCharacterIdSkillqueue200Ok struct {

	// finish_date string
	FinishDate time.Time `json:"finish_date,omitempty"`

	// finished_level integer
	FinishedLevel int32 `json:"finished_level,omitempty"`

	// level_end_sp integer
	LevelEndSp int32 `json:"level_end_sp,omitempty"`

	// Amount of SP that was in the skill when it started training it's current level. Used to calculate % of current level complete.
	LevelStartSp int32 `json:"level_start_sp,omitempty"`

	// queue_position integer
	QueuePosition int32 `json:"queue_position,omitempty"`

	// skill_id integer
	SkillId int32 `json:"skill_id,omitempty"`

	// start_date string
	StartDate time.Time `json:"start_date,omitempty"`

	// training_start_sp integer
	TrainingStartSp int32 `json:"training_start_sp,omitempty"`
}
