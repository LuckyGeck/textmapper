/*   templates.cpp
 *
 *   Lapg (Lexical Analyzer and Parser Generator)
 *   Copyright (C) 2002-04  Eugeniy Gryaznov (gryaznov@front.ru)
 *
 *   This program is free software; you can redistribute it and/or modify
 *   it under the terms of the GNU General Public License as published by
 *   the Free Software Foundation; either version 2 of the License, or
 *   (at your option) any later version.
 *
 *   This program is distributed in the hope that it will be useful,
 *   but WITHOUT ANY WARRANTY; without even the implied warranty of
 *   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *   GNU General Public License for more details.
 *
 *   You should have received a copy of the GNU General Public License
 *   along with this program; if not, write to the Free Software
 *   Foundation, Inc., 59 Temple Place, Suite 330, Boston, MA  02111-1307  USA
 */

const char *templ_cs =
	"// @target\n"
	"@nativecode\n"
	"namespace @namespace {\n"
	"\n"
	"  public class @classname {\n"
	"\t@nativecodeall\n"
	"\n"
	"${pos\n"
	"\tpublic struct lapg_place {\n"
	"$#pos1\t\tpublic int line;\n"
	"$#pos2\t\tpublic int line, column;\n"
	"$#pos3\t\tpublic int line, offset;\n"
	"\t};\n"
	"$}\n"
	"\n"
	"\tpublic struct lapg_symbol {\n"
	"\t\tpublic object sym;\n"
	"\t\tpublic int  lexem, state;\n"
	"$#pos\t\tpublic lapg_place pos;\n"
	"$#lexemend\t\tpublic lapg_place endpos;\n"
	"\t};\n"
	"\n"
	"\tstatic readonly short[] lapg_char2no = new short[] {\n"
	"\t\t@char2no\n"
	"\t};\n"
	"\n"
	"\tstatic readonly short[,] lapg_lexem = new short[@lstates,@lchars] {\n"
	"\t\t@lexem\n"
	"\t};\n"
	"\n"
	"\tstatic readonly int[] lapg_action = new int[@nstates] {\n"
	"\t\t@action\n"
	"\t};\n"
	"\n"
	"${nactions\n"
	"\tstatic readonly short[] lapg_lalr = new short[@nactions] {\n"
	"\t\t@lalr\n"
	"\t};\n"
	"$}\n"
	"\n"
	"\tstatic readonly short[] lapg_sym_goto = new short[@nsyms+1] {\n"
	"\t\t@sym_goto\n"
	"\t};\n"
	"\n"
	"\tstatic readonly short[] lapg_sym_from = new short[@gotosize] {\n"
	"\t\t@sym_from\n"
	"\t};\n"
	"\n"
	"\tstatic readonly short[] lapg_sym_to = new short[@gotosize] {\n"
	"\t\t@sym_to\n"
	"\t};\n"
	"\n"
	"\tstatic readonly short[] lapg_rlen = new short[@rules] {\n"
	"\t\t@rlen\n"
	"\t};\n"
	"\n"
	"\tstatic readonly short[] lapg_rlex = new short[@rules] {\n"
	"\t\t@rlex\n"
	"\t};\n"
	"\n"
	"\t#if DEBUG_syntax\n"
	"\tstatic readonly string[] lapg_syms = new string[] {\n"
	"\t\t@syms\n"
	"\t};\n"
	"\t#endif\n"
	"\n"
	"\tpublic enum Tokens {\n"
	"\t\t@tokenenum\n"
	"\t}\n"
	"\n"
	"\tstatic int lapg_next( int state, int symbol ) {\n"
	"${nactions\n"
	"\t\tint p;\n"
	"\t\tif( lapg_action[state] < -2 ) {\n"
	"\t\t\tfor( p = - lapg_action[state] - 3; lapg_lalr[p] >= 0; p += 2 )\n"
	"\t\t\t\tif( lapg_lalr[p] == symbol ) break;\n"
	"\t\t\treturn lapg_lalr[p+1];\n"
	"\t\t}\n"
	"$}\n"
	"\t\treturn lapg_action[state];\n"
	"\t}\n"
	"\n"
	"\tstatic int lapg_state_sym( int state, int symbol ) {\n"
	"\t\tint min = lapg_sym_goto[symbol], max = lapg_sym_goto[symbol+1]-1;\n"
	"\t\tint i, e;\n"
	"\n"
	"\t\twhile( min <= max ) {\n"
	"\t\t\te = (min + max) >> 1;\n"
	"\t\t\ti = lapg_sym_from[e];\n"
	"\t\t\tif( i == state )\n"
	"\t\t\t\treturn lapg_sym_to[e];\n"
	"\t\t\telse if( i < state )\n"
	"\t\t\t\tmin = e + 1;\n"
	"\t\t\telse\n"
	"\t\t\t\tmax = e - 1;\n"
	"\t\t}\n"
	"\t\treturn -1;\n"
	"\t}\n"
	"\n"
	"\tpublic bool parse() {\n"
	"\n"
	"\t\tbyte[]        token = new byte[@maxtoken];\n"
	"\t\tint           lapg_head = 0, group = 0, lapg_i, lapg_size, chr;\n"
	"\t\tlapg_symbol[] lapg_m = new lapg_symbol[@maxstack];\n"
	"\t\tlapg_symbol   lapg_n;\n"
	"\t\tSystem.Text.ASCIIEncoding AE = new System.Text.ASCIIEncoding();\n"
	"$#pos\t\tlapg_place    lapg_current;\n"
	"$#error\t\tint           lapg_symbols_ok = 4;\n"
	"\n"
	"$#pos1\t\tlapg_current.line = 1;\n"
	"$#pos2\t\tlapg_current.line = lapg_current.column = 1;\n"
	"$#pos3\t\tlapg_current.offset = 0;\n"
	"$#pos3\t\tlapg_current.line = 1;\n"
	"\t\tlapg_m[0].state = 0; lapg_n.lexem = lapg_n.state = 0;\n"
	"\t\t@next;\n"
	"\n"
	"\t\tdo {\n"
	"$#pos\t\t\tlapg_n.pos = lapg_current;\n"
	"\t\t\tfor( lapg_size = 0, lapg_i = group; lapg_i >= 0; ) {\n"
	"\t\t\t\tif( lapg_size < @maxtoken-1 ) token[lapg_size++] = (byte)chr;\n"
	"\t\t\t\tlapg_i = lapg_lexem[lapg_i,lapg_char2no[chr]];\n"
	"\t\t\t\tif( lapg_i >= -1 && chr != 0 ) { \n"
	"$#pos1\t\t\t\t\tif( chr == '\\n' ) lapg_current.line++;\n"
	"$#pos2\t\t\t\t\tlapg_current.column++;\n"
	"$#pos2\t\t\t\t\tif( chr == '\\n' ) { lapg_current.column = 1; lapg_current.line++; }\n"
	"$#pos3\t\t\t\t\tlapg_current.offset++;\n"
	"$#pos3\t\t\t\t\tif( chr == '\\n' ) lapg_current.line++;\n"
	"\t\t\t\t\t@next;\n"
	"\t\t\t\t}\n"
	"\t\t\t}\n"
	"$#lexemend\t\t\tlapg_n.endpos = lapg_current;\n"
	"\t\t\ttoken[lapg_size] = 0;\n"
	"\n"
	"\t\t\tif( lapg_i == -1 ) {\n"
	"$#pos0\t\t\t\terror(@errprefix System.String.Format( \"invalid lexem: `{0}`, skipped\", new System.String(AE.GetChars(token,0,lapg_size)) ) );\n"
	"$#pos1\t\t\t\terror(@errprefix System.String.Format( \"invalid lexem at line {0}: `{1}`, skipped\", lapg_n.pos.line, new System.String(AE.GetChars(token,0,lapg_size)) ) );\n"
	"$#pos2\t\t\t\terror(@errprefix System.String.Format( \"invalid lexem at line {0}, column {1}: `{2}`, skipped\", lapg_n.pos.line, lapg_n.pos.column, new System.String(AE.GetChars(token,0,lapg_size)) ) );\n"
	"$#pos3\t\t\t\terror(@errprefix System.String.Format( \"invalid lexem at line {0}: `{1}`, skipped\", lapg_n.pos.line, new System.String(AE.GetChars(token,0,lapg_size)) ) );\n"
	"\t\t\t\tcontinue;\n"
	"\t\t\t}\n"
	"\n"
	"\t\t\ttoken[lapg_size-1] = 0;\n"
	"\t\t\tlapg_n.lexem = -lapg_i-2;\n"
	"\t\t\tlapg_n.sym = null;\n"
	"\n"
	"${lexemactions\n"
	"\t\t\tswitch( lapg_n.lexem ) {\n"
	"${eachlexem\n"
	"\t\t\t\tcase @lexemnum:\n"
	"\t\t\t\t\t@lexemaction\n"
	"$}\n"
	"\t\t\t}\n"
	"$}\n"
	"\n"
	"\t\t\tskip_symbols:\n"
	"\t\t\tdo {\n"
	"\t\t\t\tlapg_i = lapg_next( lapg_m[lapg_head].state, lapg_n.lexem );\n"
	"\n"
	"\t\t\t\tif( lapg_i >= 0 ) {\n"
	"\t\t\t\t\tlapg_symbol lapg_gg;\n"
	"\t\t\t\t\tlapg_gg.sym = (lapg_rlen[lapg_i]!=0)?lapg_m[lapg_head+1-lapg_rlen[lapg_i]].sym:null;\n"
	"\t\t\t\t\tlapg_gg.lexem = lapg_rlex[lapg_i];\n"
	"\t\t\t\t\tlapg_gg.state = 0;\n"
	"\t\t\t\t\t#if DEBUG_syntax\n"
	"\t\t\t\t\t\tSystem.Console.WriteLine( \"reduce to {0}\", lapg_syms[lapg_rlex[lapg_i]] );\n"
	"\t\t\t\t\t#endif\n"
	"$#pos\t\t\t\t\tlapg_gg.pos = (lapg_rlen[lapg_i]!=0)?lapg_m[lapg_head+1-lapg_rlen[lapg_i]].pos:lapg_n.pos;\n"
	"$#lexemend\t\t\t\t\tlapg_gg.endpos = (lapg_rlen[lapg_i]!=0)?lapg_m[lapg_head].endpos:lapg_n.pos;\n"
	"${ruleactions\n"
	"\t\t\t\t\tswitch( lapg_i ) {\n"
	"${eachaction\n"
	"\t\t\t\t\t\tcase @rulenum:\n"
	"\t\t\t\t\t\t\t@ruleaction\n"
	"$}\n"
	"\t\t\t\t\t}\n"
	"$}\n"
	"\t\t\t\t\tfor( int e = lapg_rlen[lapg_i]; e > 0; e-- ) lapg_m[lapg_head--].sym = null;\n"
	"\t\t\t\t\tlapg_m[++lapg_head] = lapg_gg;\n"
	"\t\t\t\t\tlapg_m[lapg_head].state = lapg_state_sym( lapg_m[lapg_head-1].state, lapg_gg.lexem );\n"
	"\t\t\t\t} else if( lapg_i == -1 ) {\n"
	"\t\t\t\t\tlapg_m[++lapg_head] = lapg_n;\n"
	"\t\t\t\t\tlapg_m[lapg_head].state = lapg_state_sym( lapg_m[lapg_head-1].state, lapg_n.lexem );\n"
	"$#error\t\t\t\t\tlapg_symbols_ok++;\n"
	"\t\t\t\t\t#if DEBUG_syntax\n"
	"\t\t\t\t\t\tSystem.Console.WriteLine( \"shift: {0} ({1})\", lapg_syms[lapg_n.lexem], new System.String(AE.GetChars(token,0,lapg_size-1)) );\n"
	"\t\t\t\t\t#endif\n"
	"\t\t\t\t}\n"
	"\n"
	"\t\t\t} while( lapg_i >= 0 && lapg_m[lapg_head].state != -1 );\n"
	"\n"
	"\t\t\tif( (lapg_i == -2 || lapg_m[lapg_head].state == -1) && lapg_n.lexem != 0 ) {\n"
	"${error\t\t\t\t\n"
	"\t\t\t\twhile( lapg_head >= 0 && lapg_state_sym( lapg_m[lapg_head].state, @error ) == -1 ) {\n"
	"\t\t\t\t\tlapg_m[lapg_head].sym = null;\n"
	"\t\t\t\t\tlapg_head--;\n"
	"\t\t\t\t}\n"
	"\t\t\t\tif( lapg_head >= 0 ) {\n"
	"\t\t\t\t\tlapg_head++;\n"
	"\t\t\t\t\tlapg_m[lapg_head].lexem = @error;\n"
	"\t\t\t\t\tlapg_m[lapg_head].sym = null;\n"
	"\t\t\t\t\tlapg_m[lapg_head].state = lapg_state_sym( lapg_m[lapg_head-1].state, 72 );\n"
	"$#pos\t\t\t\t\tlapg_m[lapg_head].pos = lapg_n.pos;\n"
	"\t\t\t\t\tif( lapg_symbols_ok >= 4 )\n"
	"$#pos0\t\t\t\t\t\terror(@errprefix System.String.Format( \"syntax error\" ) );\n"
	"$#pos1\t\t\t\t\t\terror(@errprefix System.String.Format( \"syntax error before line {0}\", lapg_n.pos.line ) );\n"
	"$#pos2\t\t\t\t\t\terror(@errprefix System.String.Format( \"syntax error before line {0}, column {1}\", lapg_n.pos.line, lapg_n.pos.column ) );\n"
	"$#pos3\t\t\t\t\t\terror(@errprefix System.String.Format( \"syntax error before line {0}\", lapg_n.pos.line ) );\n"
	"\t\t\t\t\tlapg_i = lapg_symbols_ok;\n"
	"\t\t\t\t\tlapg_symbols_ok = 0;\n"
	"\t\t\t\t\tif( lapg_i > 1 )\n"
	"\t\t\t\t\t\tgoto skip_symbols;\n"
	"\t\t\t\t\telse\n"
	"\t\t\t\t\t\tcontinue;\n"
	"\t\t\t\t} else lapg_head = 0;\n"
	"$}\n"
	"\t\t\t\tbreak;\n"
	"\t\t\t}\n"
	"\n"
	"\t\t} while( lapg_n.lexem != 0 );\n"
	"\n"
	"${noterror\n"
	"\t\tif( lapg_m[lapg_head].state != @nstates-1 ) {\n"
	"$}\n"
	"${error\n"
	"\t\tif( lapg_m[lapg_head].state != @nstates-1 && lapg_symbols_ok >= 4 ) {\n"
	"$}\n"
	"$#pos0\t\t\terror(@errprefix System.String.Format( \"syntax error\" ) );\n"
	"$#pos1\t\t\terror(@errprefix System.String.Format( \"syntax error before line {0}\", lapg_n.pos.line ) );\n"
	"$#pos2\t\t\terror(@errprefix System.String.Format( \"syntax error before line {0}, column {1}\", lapg_n.pos.line, lapg_n.pos.column ) );\n"
	"$#pos3\t\t\terror(@errprefix System.String.Format( \"syntax error before line {0}\", lapg_n.pos.line ) );\n"
	"\t\t\treturn false;\n"
	"\t\t};\n"
	"\t\treturn true;\n"
	"\t}\n"
	"  }\n"
	"}\n"
;

const char *templ_cpp =
	"// @target\n"
	"@nativecodeall\n"
	"${pos\n"
	"struct lapg_place {\n"
	"$#pos1\tint line;\n"
	"$#pos2\tint line, column;\n"
	"$#pos3\tint line, offset;\n"
	"};\n"
	"$}\n"
	"\n"
	"struct lapg_symbol {\n"
	"\tvoid *sym;\n"
	"\tint  lexem, state;\n"
	"$#pos\tstruct lapg_place pos;\n"
	"$#lexemend\tstruct lapg_place endpos;\n"
	"};\n"
	"\n"
	"static const char lapg_char2no[256] = {\n"
	"@char2no\n"
	"};\n"
	"\n"
	"static const short lapg_lexem[@lstates][@lchars] = {\n"
	"@lexem\n"
	"};\n"
	"\n"
	"static const int lapg_action[@nstates] = {\n"
	"@action\n"
	"};\n"
	"\n"
	"${nactions\n"
	"static const short lapg_lalr[@nactions] = {\n"
	"@lalr\n"
	"};\n"
	"$}\n"
	"\n"
	"static const short lapg_sym_goto[@nsyms+1] = {\n"
	"@sym_goto\n"
	"};\n"
	"\n"
	"static const short lapg_sym_from[@gotosize] = {\n"
	"@sym_from\n"
	"};\n"
	"\n"
	"static const short lapg_sym_to[@gotosize] = {\n"
	"@sym_to\n"
	"};\n"
	"\n"
	"static const short lapg_rlen[@rules] = {\n"
	"@rlen\n"
	"};\n"
	"\n"
	"static const short lapg_rlex[@rules] = {\n"
	"@rlex\n"
	"};\n"
	"\n"
	"#ifdef DEBUG_syntax\n"
	"static const char *lapg_syms[] = {\n"
	"@syms\n"
	"};\n"
	"#endif\n"
	"\n"
	"static inline int lapg_next( int state, int symbol )\n"
	"{\n"
	"${nactions\n"
	"\tif( lapg_action[state] < -2 ) {\n"
	"\t\tconst short *p = lapg_lalr - lapg_action[state] - 3;\n"
	"\t\tfor( ;*p >= 0; p += 2 )\n"
	"\t\t\tif( *p == symbol ) break;\n"
	"\t\treturn p[1];\n"
	"\t}\n"
	"$}\n"
	"\treturn lapg_action[state];\n"
	"}\n"
	"\n"
	"static inline int lapg_state_sym( int state, int symbol )\n"
	"{\n"
	"\tint min = lapg_sym_goto[symbol], max = lapg_sym_goto[symbol+1]-1;\n"
	"\tint i, e;\n"
	"\n"
	"\twhile( min <= max ) {\n"
	"\t\te = (min + max) >> 1;\n"
	"\t\ti = lapg_sym_from[e];\n"
	"\t\tif( i == state )\n"
	"\t\t\treturn lapg_sym_to[e];\n"
	"\t\telse if( i < state )\n"
	"\t\t\tmin = e + 1;\n"
	"\t\telse\n"
	"\t\t\tmax = e - 1;\n"
	"\t}\n"
	"\treturn -1;\n"
	"}\n"
	"\n"
	"int @classname::parse()\n"
	"{\n"
	"\tchar *token = new char[@maxtoken];\n"
	"\tint  lapg_head = 0, group = 0, lapg_i, lapg_size, chr;\n"
	"\tlapg_symbol *lapg_m = new lapg_symbol[@maxstack];\n"
	"\tlapg_symbol lapg_n = { NULL, -1, 0 };\n"
	"$#pos3\tlapg_place lapg_current = { 1, 0 };\n"
	"$#pos2\tlapg_place lapg_current = { 1, 1 };\n"
	"$#pos1\tlapg_place lapg_current = { 1 };\n"
	"\n"
	"$#error    int lapg_symbols_ok = 4;\n"
	"\tlapg_m[0].state = 0;\n"
	"\t@next;\n"
	"\n"
	"\tdo {\n"
	"$#pos\t\tlapg_n.pos = lapg_current;\n"
	"\t\tfor( lapg_size = 0, lapg_i = group; lapg_i >= 0; ) {\n"
	"\t\t\tif( lapg_size < @maxtoken-1 ) token[lapg_size++] = chr;\n"
	"\t\t\tlapg_i = lapg_lexem[lapg_i][lapg_char2no[chr]];\n"
	"\t\t\tif( lapg_i >= -1 && chr ) { \n"
	"$#pos1\t\t\t\tif( chr == '\\n' ) lapg_current.line++;\n"
	"$#pos2\t\t\t\tlapg_current.column++;\n"
	"$#pos2\t\t\t\tif( chr == '\\n' ) lapg_current.column = 1, lapg_current.line++;\n"
	"$#pos3\t\t\t\tlapg_current.offset++;\n"
	"$#pos3\t\t\t\tif( chr == '\\n' ) lapg_current.line++;\n"
	"\t\t\t\t@next;\n"
	"\t\t\t}\n"
	"\t\t}\n"
	"$#lexemend\t\tlapg_n.endpos = lapg_current;\n"
	"\t\ttoken[lapg_size] = 0;\n"
	"\n"
	"\t\tif( lapg_i == -1 ) {\n"
	"$#pos0\t\t\terror( @errprefix\"invalid lexem: `%s`, skipped\\n\", token );\n"
	"$#pos1\t\t\terror( @errprefix\"invalid lexem at line %i: `%s`, skipped\\n\", lapg_n.pos.line, token );\n"
	"$#pos2\t\t\terror( @errprefix\"invalid lexem at line %i, column %i: `%s`, skipped\\n\", lapg_n.pos.line, lapg_n.pos.column, token );\n"
	"$#pos3\t\t\terror( @errprefix\"invalid lexem at line %i: `%s`, skipped\\n\", lapg_n.pos.line, token );\n"
	"\t\t\tcontinue;\n"
	"\t\t}\n"
	"\n"
	"\t\ttoken[lapg_size-1] = 0;\n"
	"\t\tlapg_n.lexem = -lapg_i-2;\n"
	"\t\tlapg_n.sym = NULL;\n"
	"${lexemactions\n"
	"\t\tswitch( lapg_n.lexem ) {\n"
	"${eachlexem\n"
	"\t\t\tcase @lexemnum: {\n"
	"\t\t\t\t@lexemactioncpp\n"
	"\t\t\t} break;\n"
	"$}\n"
	"\t\t}\n"
	"$}\n"
	"\n"
	"\t\tskip_symbols:\n"
	"\t\tdo {\n"
	"\t\t\tlapg_i = lapg_next( lapg_m[lapg_head].state, lapg_n.lexem );\n"
	"\n"
	"\t\t\tif( lapg_i >= 0 ) {\n"
	"\t\t\t\tlapg_symbol lapg_gg={(lapg_rlen[lapg_i])?lapg_m[lapg_head+1-lapg_rlen[lapg_i]].sym:NULL,lapg_rlex[lapg_i],0 };\n"
	"\t\t\t\t#ifdef DEBUG_syntax\n"
	"\t\t\t\t\tfprintf( stdout, \"reduce to %s\\n\", lapg_syms[lapg_rlex[lapg_i]] );\n"
	"\t\t\t\t#endif\n"
	"$#pos\t\t\t\tlapg_gg.pos = (lapg_rlen[lapg_i])?lapg_m[lapg_head+1-lapg_rlen[lapg_i]].pos:lapg_n.pos;\n"
	"$#lexemend\t\t\t\tlapg_gg.endpos = (lapg_rlen[lapg_i])?lapg_m[lapg_head].endpos:lapg_n.pos;\n"
	"${ruleactions\n"
	"\t\t\t\tswitch( lapg_i ) {\n"
	"${eachaction\n"
	"\t\t\t\t\tcase @rulenum: {\n"
	"\t\t\t\t\t\t@ruleactioncpp\n"
	"\t\t\t\t\t} break;\n"
	"$}\n"
	"\t\t\t\t}\n"
	"$}\n"
	"\t\t\t\tlapg_head -= lapg_rlen[lapg_i];\n"
	"\t\t\t\tlapg_m[++lapg_head] = lapg_gg;\n"
	"\t\t\t\tlapg_m[lapg_head].state = lapg_state_sym( lapg_m[lapg_head-1].state, lapg_gg.lexem );\n"
	"\t\t\t} else if( lapg_i == -1 ) {\n"
	"\t\t\t\tlapg_m[++lapg_head] = lapg_n;\n"
	"\t\t\t\tlapg_m[lapg_head].state = lapg_state_sym( lapg_m[lapg_head-1].state, lapg_n.lexem );\n"
	"$#error\t\t\t\tlapg_symbols_ok++;\n"
	"\t\t\t\t#ifdef DEBUG_syntax\n"
	"\t\t\t\t\tfprintf( stdout, \"shift: %s (%s)\\n\", lapg_syms[lapg_n.lexem], token );\n"
	"\t\t\t\t#endif\n"
	"\t\t\t}\n"
	"\n"
	"\t\t} while( lapg_i >= 0 && lapg_m[lapg_head].state != -1 );\n"
	"\n"
	"\t\tif( (lapg_i == -2 || lapg_m[lapg_head].state == -1) && lapg_n.lexem != 0 ) {\n"
	"${error\n"
	"\t\t\twhile( lapg_head >= 0 && lapg_state_sym( lapg_m[lapg_head].state, @error ) == -1 )\n"
	"\t\t\t\tlapg_head--;\n"
	"\n"
	"\t\t\tif( lapg_head >= 0 ) {\n"
	"\t\t\t\tlapg_head++;\n"
	"\t\t\t\tlapg_m[lapg_head].lexem = @error;\n"
	"\t\t\t\tlapg_m[lapg_head].sym = NULL;\n"
	"\t\t\t\tlapg_m[lapg_head].state = lapg_state_sym( lapg_m[lapg_head-1].state, @error );\n"
	"$#pos\t\t\t\tlapg_m[lapg_head].pos = lapg_n.pos;\n"
	"\t\t\t\tif( lapg_symbols_ok >= 4 )\n"
	"$#pos0\t\t\t\t\terror( @errprefix\"syntax error\\n\" );\n"
	"$#pos1\t\t\t\t\terror( @errprefix\"syntax error before line %i\\n\", lapg_n.pos.line );\n"
	"$#pos2\t\t\t\t\terror( @errprefix\"syntax error before line %i, column %i\\n\", lapg_n.pos.line, lapg_n.pos.column );\n"
	"$#pos3\t\t\t\t\terror( @errprefix\"syntax error before line %i\\n\", lapg_n.pos.line );\n"
	"\t\t\t\tlapg_i = lapg_symbols_ok;\n"
	"\t\t\t\tlapg_symbols_ok = 0;\n"
	"\t\t\t\tif( lapg_i > 1 )\n"
	"\t\t\t\t\tgoto skip_symbols;\n"
	"\t\t\t\telse\n"
	"\t\t\t\t\tcontinue;\n"
	"\t\t\t} else lapg_head = 0;\n"
	"$}\n"
	"\t\t\tbreak;\n"
	"\t\t}\n"
	"\n"
	"\t} while( lapg_n.lexem );\n"
	"\n"
	"\tif( lapg_m[lapg_head].state == @nstates-1 ) lapg_i = 1; else lapg_i = 0;\n"
	"\tdelete[] lapg_m;\n"
	"\tdelete[] token;\n"
	"\n"
	"${noterror\n"
	"$#pos0\tif( !lapg_i ) error( @errprefix\"syntax error\\n\" );\n"
	"$#pos1\tif( !lapg_i ) error( @errprefix\"syntax error before line %i\\n\", lapg_n.pos.line );\n"
	"$#pos2\tif( !lapg_i ) error( @errprefix\"syntax error before line %i, column %i\\n\", lapg_n.pos.line, lapg_n.pos.column );\n"
	"$#pos3\tif( !lapg_i ) error( @errprefix\"syntax error before line %i\\n\", lapg_n.pos.line );\n"
	"$}\n"
	"${error\n"
	"$#pos0\tif( !lapg_i && lapg_symbols_ok >= 4 ) error( @errprefix\"syntax error\\n\" );\n"
	"$#pos1\tif( !lapg_i && lapg_symbols_ok >= 4 ) error( @errprefix\"syntax error before line %i\\n\", lapg_n.pos.line );\n"
	"$#pos2\tif( !lapg_i && lapg_symbols_ok >= 4 ) error( @errprefix\"syntax error before line %i, column %i\\n\", lapg_n.pos.line, lapg_n.pos.column );\n"
	"$#pos3\tif( !lapg_i && lapg_symbols_ok >= 4 ) error( @errprefix\"syntax error before line %i\\n\", lapg_n.pos.line );\n"
	"$}\n"
	"\treturn lapg_i;\n"
	"}\n"
;

