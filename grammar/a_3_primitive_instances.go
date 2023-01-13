package grammar

//
// A.3 Primitive instances
//

//
// A.3.1 Primitive instantiation and instances
//

/*
 * gate_instantiation ::=
 *   cmos_switchtype [delay3] cmos_switch_instance { , cmos_switch_instance } ;
 *   | enable_gatetype [drive_strength] [delay3] enable_gate_instance { , enable_gate_instance } ;
 *   | mos_switchtype [delay3] mos_switch_instance { , mos_switch_instance } ;
 *   | n_input_gatetype [drive_strength] [delay2] n_input_gate_instance { , n_input_gate_instance } ;
 *   | n_output_gatetype [drive_strength] [delay2] n_output_gate_instance
 *   { , n_output_gate_instance } ;
 *   | pass_en_switchtype [delay2] pass_enable_switch_instance { , pass_enable_switch_instance } ;
 *   | pass_switchtype pass_switch_instance { , pass_switch_instance } ;
 *   | pulldown [pulldown_strength] pull_gate_instance { , pull_gate_instance } ;
 *   | pullup [pullup_strength] pull_gate_instance { , pull_gate_instance } ;
 */

/*
 * cmos_switch_instance ::= [ name_of_instance ] ( output_terminal , input_terminal ,
 *   ncontrol_terminal , pcontrol_terminal )
 */

/*
 * enable_gate_instance ::= [ name_of_instance ] ( output_terminal , input_terminal , enable_terminal )
 */

/*
 * mos_switch_instance ::= [ name_of_instance ] ( output_terminal , input_terminal , enable_terminal )
 */

/*
 * n_input_gate_instance ::= [ name_of_instance ] ( output_terminal , input_terminal { , input_terminal } )
 */

/*
 * n_output_gate_instance ::= [ name_of_instance ] ( output_terminal { , output_terminal } ,
 *   input_terminal )
 */

/*
 * pass_switch_instance ::= [ name_of_instance ] ( inout_terminal , inout_terminal )
 */

/*
 * pass_enable_switch_instance ::= [ name_of_instance ] ( inout_terminal , inout_terminal ,
 *   enable_terminal )
 */

/*
 * pull_gate_instance ::= [ name_of_instance ] ( output_terminal )
 */

//
// A.3.2 Primitive strengths
//

/*
 * pulldown_strength ::=
 *   ( strength0 , strength1 )
 *   | ( strength1 , strength0 )
 *   | ( strength0 )
 */

/*
 * pullup_strength ::=
 *   ( strength0 , strength1 )
 *   | ( strength1 , strength0 )
 *   | ( strength1 )
 */

//
// A.3.3 Primitive terminals
//

/*
 * enable_terminal ::= expression
 */

/*
 * inout_terminal ::= net_lvalue
 */

/*
 * input_terminal ::= expression
 */

/*
 * ncontrol_terminal ::= expression
 */

/*
 * output_terminal ::= net_lvalue
 */

/*
 * pcontrol_terminal ::= expression
 */

//
// A.3.4 Primitive gate and switch types
//

/*
 * cmos_switchtype ::= cmos | rcmos
 */

/*
 * enable_gatetype ::= bufif0 | bufif1 | notif0 | notif1
 */

/*
 * mos_switchtype ::= nmos | pmos | rnmos | rpmos
 */

/*
 * n_input_gatetype ::= and | nand | or | nor | xor | xnor
 */

/*
 * n_output_gatetype ::= buf | not
 */

/*
 * pass_en_switchtype ::= tranif0 | tranif1 | rtranif1 | rtranif0
 */

/*
 * pass_switchtype ::= tran | rtran
 */
