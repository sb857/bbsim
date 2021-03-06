/*
 * Copyright (c) 2018 - present.  Boling Consulting Solutions (bcsw.net)
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 * http://www.apache.org/licenses/LICENSE-2.0
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
/*
 * NOTE: This file was generated, manual edits will be overwritten!
 *
 * Generated by 'goCodeGenerator.py':
 *              https://github.com/cboling/OMCI-parser/README.md
 */
package generated

import "github.com/deckarep/golang-set"

const PhysicalPathTerminationPointCesUniClassId ClassID = ClassID(12)

var physicalpathterminationpointcesuniBME *ManagedEntityDefinition

// PhysicalPathTerminationPointCesUni (class ID #12)
//	This ME represents the point at a CES UNI in the ONU where the physical path terminates and
//	physical level functions are performed.
//
//	The ONU automatically creates an instance of this ME per port:
//
//	•	when the ONU has CES ports built into its factory configuration;
//
//	•	when a cardholder is provisioned to expect a circuit pack of a CES type;
//
//	•	when a cardholder provisioned for plug-and-play is equipped with a circuit pack of a CES type.
//	Note that the installation of a plug-and-play card may indicate the presence of CES ports via
//	equipment ID as well as its type and indeed may cause the ONU to instantiate a port-mapping
//	package that specifies CES ports.
//
//	The ONU automatically deletes instances of this ME when a cardholder is neither provisioned to
//	expect a CES circuit pack, nor is it equipped with a CES circuit pack.
//
//	Relationships
//		An instance of this ME is associated with each real or pre-provisioned CES port. It can be
//		linked from a GEM IW TP, a pseudowire TP or a logical N × 64 kbit/s CTP.
//
//	Attributes
//		Managed Entity Id
//			Managed entity ID: This attribute uniquely identifies each instance of this ME. This 2 byte
//			number indicates the physical position of the UNI. The first byte is the slot ID (defined in
//			clause 9.1.5). The second byte is the port ID, with the range 1..255. (R) (mandatory) (2 bytes)
//
//		Expected Type
//			Upon ME instantiation, the ONU sets this attribute to 0. (R, W) (mandatory) (1 byte)
//
//		Sensed Type
//			Sensed type:	If the value of expected type is not 0, then the value of sensed type equals the
//			value of expected type. If expected type = 0, then the value of sensed type is one of the
//			compatible values from Table 9.1.5-1. Upon ME instantiation, the ONU sets this attribute to 0 or
//			to the value that reflects the physically present equipment. (R) (mandatory if the ONU supports
//			circuit packs with configurable interface types, e.g., C1.5/2/6.3) (1 byte)
//
//		Ces Loopback Configuration
//			Upon ME instantiation, the ONU sets this attribute to 0. (R, W) (mandatory) (1 byte)
//
//		Administrative State
//			Administrative state: This attribute locks (1) and unlocks (0) the functions performed by this
//			ME. Administrative state is further described in clause A.1.6. (R, W) (mandatory) (1 byte)
//
//		Operational State
//			Operational state: This attribute indicates whether the ME is capable of performing its
//			function. Valid values are enabled (0) and disabled (1). (R) (optional) (1 byte)
//
//		Framing
//			Upon ME instantiation, the ONU sets this attribute to a value that reflects the vendor's
//			default. (R, W) (optional) (1 byte)
//
//		Encoding
//			Upon ME instantiation, the ONU sets this attribute to 0. (R, W) (mandatory for DS1 and DS3
//			interfaces) (1 byte)
//
//		Line Length
//			Line length:	This attribute specifies the length of the twisted pair cable from a DS1 physical
//			UNI to the DSX-1 cross-connect point or the length of coaxial cable from a DS3 physical UNI to
//			the DSX-3 cross-connect point. Valid values are given in Table 9.8.1-1. Upon ME instantiation
//			for a DS1 interface, the ONU assigns the value 0 for non-power feed type DS1 and the value 6 for
//			power feed type DS1. Upon ME instantiation for a DS3 interface, the ONU sets this attribute to
//			0x0F. (R, W) (optional) (1 byte)
//
//		Ds1 Mode
//			In the event of conflicting values between this attribute and the (also optional) line length
//			attribute, the line length attribute is taken to be valid. This permits the separation of line
//			build-out (LBO) and power settings from smart jack and FDL behaviour. Upon ME instantiation, the
//			ONU sets this attribute to 0. (R, W) (optional) (1 byte)
//
//		Arc
//			ARC:	See clause A.1.4.3. (R, W) (optional) (1 byte)
//
//		Arc Interval
//			ARC interval: See clause A.1.4.3. (R, W) (optional) (1 byte)
//
//		Line Type
//			(R, W) (mandatory for DS3, E3 and multi-configuration interfaces, not applicable to other
//			interfaces) (1 byte)
//
type PhysicalPathTerminationPointCesUni struct {
	ManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	physicalpathterminationpointcesuniBME = &ManagedEntityDefinition{
		Name:    "PhysicalPathTerminationPointCesUni",
		ClassID: 12,
		MessageTypes: mapset.NewSetWith(
			Get,
			Set,
		),
		AllowedAttributeMask: 0XFFF0,
		AttributeDefinitions: AttributeDefinitionMap{
			0:  Uint16Field("ManagedEntityId", 0, mapset.NewSetWith(Read), false, false, false, false, 0),
			1:  ByteField("ExpectedType", 0, mapset.NewSetWith(Read, Write), false, false, false, false, 1),
			2:  ByteField("SensedType", 0, mapset.NewSetWith(Read), true, false, false, false, 2),
			3:  ByteField("CesLoopbackConfiguration", 0, mapset.NewSetWith(Read, Write), true, false, false, false, 3),
			4:  ByteField("AdministrativeState", 0, mapset.NewSetWith(Read, Write), false, false, false, false, 4),
			5:  ByteField("OperationalState", 0, mapset.NewSetWith(Read), true, false, true, false, 5),
			6:  ByteField("Framing", 0, mapset.NewSetWith(Read, Write), false, false, true, false, 6),
			7:  ByteField("Encoding", 0, mapset.NewSetWith(Read, Write), false, false, false, false, 7),
			8:  ByteField("LineLength", 0, mapset.NewSetWith(Read, Write), false, false, true, false, 8),
			9:  ByteField("Ds1Mode", 0, mapset.NewSetWith(Read, Write), false, false, true, false, 9),
			10: ByteField("Arc", 0, mapset.NewSetWith(Read, Write), true, false, true, false, 10),
			11: ByteField("ArcInterval", 0, mapset.NewSetWith(Read, Write), false, false, true, false, 11),
			12: ByteField("LineType", 0, mapset.NewSetWith(Read, Write), false, false, false, false, 12),
		},
	}
}

// NewPhysicalPathTerminationPointCesUni (class ID 12 creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from the wire, about to be sent on the wire.
func NewPhysicalPathTerminationPointCesUni(params ...ParamData) (*ManagedEntity, OmciErrors) {
	return NewManagedEntity(physicalpathterminationpointcesuniBME, params...)
}
